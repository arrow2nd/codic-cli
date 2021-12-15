package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = newRootCmd()

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:          "codic",
		Short:        "Unofficial client of \"codic\"",
		Long:         "Unofficial client of the naming dictionary service \"codic\"",
		Version:      Version,
		SilenceUsage: true,
	}

	root.AddCommand(
		newGetCmd(),
	)

	return root
}

// Execute 実行
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// ファイル設定
	viper.AddConfigPath(home)
	viper.SetConfigName(".codic-cli")
	viper.SetConfigType("yaml")

	// デフォルト値
	viper.SetDefault("APIKey", "")
	viper.SetDefault("Casing", "")
	viper.SetDefault("AcronymStyle", LiteralStyle)

	// ファイルが存在しない場合作成
	viper.SafeWriteConfig()

	// 設定ファイル読み込み
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
