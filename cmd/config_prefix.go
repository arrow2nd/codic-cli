package cmd

import (
	"github.com/arrow2nd/godic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var prefixList = []selectItem{
	{
		Value:   godic.MSPrefix,
		Example: "IOException",
	},
	{
		Value:   godic.CamelStrictPrefix,
		Example: "IoException",
	},
	{
		Value:   godic.LiteralPrefix,
		Example: "(No conversion)",
	},
}

func (c *Cmd) newConfigPrefixCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "prefix",
		Short: "Set the acronym type",
		Args:  cobra.NoArgs,
		RunE:  c.execConfigPrefixCmd,
	}
}

func (c *Cmd) execConfigPrefixCmd(cmd *cobra.Command, args []string) error {
	result := selectList("Prefix type", prefixList)
	if result == "" {
		return nil
	}

	viper.Set("acronymstyle", result)
	if err := viper.WriteConfig(); err != nil {
		return err
	}

	showSuccessMessage("Saved!")
	return nil
}
