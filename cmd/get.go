package cmd

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

func (c *Cmd) newGetCmd() *cobra.Command {
	new := &cobra.Command{
		Use:     "get [japanese text]",
		Short:   "Get the naming",
		Example: "  codic get ファイルを作成する",
		Args:    cobra.ExactArgs(1),
		RunE:    c.execGetCmd,
	}

	return new
}

func (c *Cmd) execGetCmd(cmd *cobra.Command, args []string) error {
	text := args[0]

	v := url.Values{}
	v.Add("text", text)

	results, err := c.api.GetTranslate(v)
	if err != nil {
		return err
	}

	fmt.Println(results[0].TranslatedText)

	return nil
}
