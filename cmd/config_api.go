package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (c *Cmd) newConfigAPICmd() *cobra.Command {
	new := &cobra.Command{
		Use:     "config [japanese text]",
		Short:   "Operate the configuration",
		Example: "  config",
		Args:    cobra.ExactArgs(1),
		RunE:    c.execConfigAPICmd,
	}

	return new
}

func (c *Cmd) execConfigAPICmd(cmd *cobra.Command, args []string) error {
	apiKey := args[0]

	viper.Set("APIKey", apiKey)

	return nil
}
