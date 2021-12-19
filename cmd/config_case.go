package cmd

import (
	"github.com/arrow2nd/codic-cli/list"
	"github.com/arrow2nd/godic"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var caseList = list.List{
	{
		Name:    "pascal",
		Value:   godic.PascalCase,
		Example: "PascalCase",
	},
	{
		Name:    "camel",
		Value:   godic.CamelCase,
		Example: "camelCase",
	},
	{
		Name:    "snake",
		Value:   godic.SnakeCase,
		Example: "snake_case",
	},
	{
		Name:    "screaming",
		Value:   godic.ScreamingSnakeCase,
		Example: "SNAKE_CASE",
	},
	{
		Name:    "kebab",
		Value:   godic.KebabCase,
		Example: "kebab-case",
	},
	{
		Name:    "space",
		Value:   "",
		Example: "space case",
	},
}

func (c *Cmd) newConfigCaseCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "case",
		Short: "Set the case type",
		Args:  cobra.NoArgs,
		RunE:  c.execConfigCaseCmd,
	}
}

func (c *Cmd) execConfigCaseCmd(cmd *cobra.Command, args []string) error {
	result := caseList.SelectList("Case type")
	if result == "" {
		return nil
	}

	viper.Set("case", result)
	if err := viper.WriteConfig(); err != nil {
		return err
	}

	showSuccessMessage("Saved!")
	return nil
}
