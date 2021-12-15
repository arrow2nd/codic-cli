package cmd

import "github.com/spf13/cobra"

func newGetCmd() *cobra.Command {
	new := &cobra.Command{
		Use:     "get",
		Short:   "Get the naming",
		Example: "codic get ファイルを作成する",
		RunE:    execGetCmd,
	}

	return new
}

func execGetCmd(cmd *cobra.Command, args []string) error {
	return nil
}
