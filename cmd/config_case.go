package cmd

import (
	"github.com/arrow2nd/godic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var caseList = []selectItem{
	{
		Value:   godic.PascalCase,
		Example: "PascalCase",
	},
	{
		Value:   godic.CamelCase,
		Example: "camelCase",
	},
	{
		Value:   godic.SnakeCase,
		Example: "snake_case",
	},
	{
		Value:   godic.ScreamingSnakeCase,
		Example: "SNAKE_CASE",
	},
	{
		Value:   godic.KebabCase,
		Example: "kabab-case",
	},
}

func (c *Cmd) newConfigCaseCmd() *cobra.Command {
	caseCmd := &cobra.Command{
		Use:   "case",
		Short: "Set the case type",
		Args:  cobra.NoArgs,
		RunE:  c.execConfigCaseCmd,
	}

	return caseCmd
}

func (c *Cmd) execConfigCaseCmd(cmd *cobra.Command, args []string) error {
	result := selectList("Case type", caseList)
	if result == "" {
		return nil
	}

	viper.Set("casing", result)
	if err := viper.WriteConfig(); err != nil {
		return err
	}

	showSuccessMessage("complete")
	return nil
}
