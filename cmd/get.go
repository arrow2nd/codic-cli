package cmd

import (
	"fmt"

	"github.com/arrow2nd/godic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (c *Cmd) newGetCmd() *cobra.Command {
	getCmd := &cobra.Command{
		Use:     "get [japanese text]",
		Short:   "Get the naming",
		Example: "  codic get ファイルを作成する",
		Args:    cobra.ExactArgs(1),
		RunE:    c.execGetCmd,
	}

	getCmd.Flags().StringP("case", "c", viper.GetString("case"), "specify the case type")
	getCmd.Flags().StringP("prefix", "p", viper.GetString("prefix"), "specify the prefix style")

	return getCmd
}

func (c *Cmd) execGetCmd(cmd *cobra.Command, args []string) error {
	text := args[0]
	caseType, _ := cmd.Flags().GetString("case")
	prefixStyle, _ := cmd.Flags().GetString("prefix")

	caseType, err := caseList.GetValue(caseType)
	if err != nil {
		return fmt.Errorf("case type is wrong, please specify one of [%s]", caseList.GetListItemsString())
	}

	prefixStyle, err = prefixList.GetValue(prefixStyle)
	if err != nil {
		return fmt.Errorf("prefix style is wrong, please specify one of [%s]", prefixList.GetListItemsString())
	}

	v := godic.CreateTranslateParam(text)
	v.Add("casing", caseType)
	v.Add("acronym_style", prefixStyle)

	results, err := c.api.GetTranslate(v)
	if err != nil {
		return err
	}

	fmt.Println(results[0].TranslatedText)

	return nil
}
