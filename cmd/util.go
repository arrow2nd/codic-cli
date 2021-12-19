package cmd

import (
	"fmt"

	"github.com/fatih/color"
)

// showSuccessMessage 完了メッセージ
func showSuccessMessage(text string) {
	fmt.Printf("%s %s\n", color.GreenString("✔"), text)
}
