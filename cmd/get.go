package cmd

import (
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	new := &cobra.Command{
		Use:     "get [japanese text]",
		Short:   "Get the naming",
		Example: "  codic get ファイルを作成する",
		Args:    cobra.ExactArgs(1),
		RunE:    execGetCmd,
	}

	return new
}

func execGetCmd(cmd *cobra.Command, args []string) error {
	// text := args[0]

	return nil
}
