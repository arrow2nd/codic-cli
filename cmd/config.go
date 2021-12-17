package cmd

import "github.com/spf13/cobra"

func (c *Cmd) newConfigCmd() *cobra.Command {
	new := &cobra.Command{
		Use:     "config [japanese text]",
		Short:   "Operate the configuration",
		Example: "  config",
	}

	new.AddCommand(c.newConfigAPICmd())

	return new
}
