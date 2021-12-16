package cmd

import "github.com/spf13/cobra"

func newConfigCmd() *cobra.Command {
	new := &cobra.Command{
		Use:     "config [japanese text]",
		Short:   "Operate the configuration",
		Example: "  config",
	}

	return new
}
