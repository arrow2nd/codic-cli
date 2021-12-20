package cmd

import "github.com/spf13/cobra"

func (c *Cmd) newConfigCmd() *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Set various parameters",
	}

	configCmd.AddCommand(
		c.newConfigAPICmd(),
		c.newConfigCaseCmd(),
		c.newConfigPrefixCmd(),
		c.newConfigListCmd(),
	)

	return configCmd
}
