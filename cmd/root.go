package cmd

import (
	"os"

	"github.com/arrow2nd/godic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// ファイル設定
	viper.AddConfigPath(home)
	viper.SetConfigName(".codic-cli")
	viper.SetConfigType("yaml")

	// デフォルト値
	viper.SetDefault("apikey", "")
	viper.SetDefault("casing", "")
	viper.SetDefault("acronymstyle", godic.LiteralPrefix)

	// ファイルが存在しない場合作成
	viper.SafeWriteConfig()

	// 設定ファイル読み込み
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

// Cmd コマンド
type Cmd struct {
	root *cobra.Command
	api  *godic.Godic
}

// New 生成
func New() *Cmd {
	newCmd := &Cmd{
		root: &cobra.Command{
			Use:          "codic",
			Short:        "Unofficial client of \"codic\"",
			Long:         "Unofficial client of the naming dictionary service \"codic\"",
			Version:      Version,
			SilenceUsage: true,
		},
		api: godic.NewClient(viper.GetString("apikey")),
	}

	newCmd.root.AddCommand(
		newCmd.newGetCmd(),
		newCmd.newConfigCmd(),
	)

	return newCmd
}

// Execute 実行
func (c *Cmd) Execute() {
	c.root.Execute()
}
