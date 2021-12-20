package cmd

import (
	"github.com/arrow2nd/codic-cli/list"
	"github.com/arrow2nd/godic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var prefixList = list.List{
	{
		Name:    "microsoft",
		Value:   godic.MSPrefix,
		Example: "IOException",
	},
	{
		Name:    "camel",
		Value:   godic.CamelStrictPrefix,
		Example: "IoException",
	},
	{
		Name:    "literal",
		Value:   godic.LiteralPrefix,
		Example: "(No conversion)",
	},
}

func (c *Cmd) newConfigPrefixCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "prefix",
		Short: "Set the default prefix style",
		Args:  cobra.NoArgs,
		RunE:  c.execConfigPrefixCmd,
	}
}

func (c *Cmd) execConfigPrefixCmd(cmd *cobra.Command, args []string) error {
	result := prefixList.SelectList("Default Prefix Style?")
	if result == "" {
		return nil
	}

	viper.Set("plefix", result)
	if err := viper.WriteConfig(); err != nil {
		return err
	}

	showSuccessMessage("Saved!")
	return nil
}
