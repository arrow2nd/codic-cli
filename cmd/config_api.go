package cmd

import (
	"fmt"

	"github.com/arrow2nd/godic"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (c *Cmd) newConfigAPICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Register your API key",
		Args:  cobra.NoArgs,
		RunE:  c.execConfigAPICmd,
	}
}

func (c *Cmd) execConfigAPICmd(cmd *cobra.Command, args []string) error {
	fmt.Println("Enter the API key for codic.")
	fmt.Println("If you don't have it, you can get it here. [ https://codic.jp/my/api_status ]")

	prompt := promptui.Prompt{
		Label:       "API Key",
		HideEntered: true,
		Validate: func(s string) error {
			if len(s) == 0 {
				return fmt.Errorf("too short")
			}
			return nil
		},
	}

	apiKey, err := prompt.Run()
	if err != nil {
		return nil
	}

	g := godic.NewClient(apiKey)
	params := godic.CreateTranslateParam("こんにちは世界")

	// 有効なAPIキーかチェック
	if _, err := g.GetTranslate(params); err != nil {
		return fmt.Errorf("not a valid API key")
	}

	viper.Set("apikey", apiKey)
	if err := viper.WriteConfig(); err != nil {
		return err
	}

	showSuccessMessage("Authentication completed!")
	return nil
}
