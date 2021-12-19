package cmd

import "github.com/spf13/cobra"

func (c *Cmd) newConfigCmd() *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Operate the configuration",
	}

	configCmd.AddCommand(
		c.newConfigAPICmd(),
		c.newConfigCaseCmd(),
	)

	return configCmd
}
