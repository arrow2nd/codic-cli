package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

// showSuccessMessage 完了メッセージ
func showSuccessMessage(text string) {
	fmt.Printf("%s %s\n", color.GreenString("✔"), text)
}

// selectItem リストアイテム
type selectItem struct {
	Value   string
	Example string
}

// selectList リスト内から1つ選択
func selectList(label string, items []selectItem) string {
	templates := &promptui.SelectTemplates{
		Label:    "{{ .Value }}?",
		Active:   `{{ ">" | cyan }} {{ .Value | cyan }}`,
		Inactive: "  {{ .Value }}",
		Selected: `{{ "` + label + `:" | faint }} {{ .Value }}`,
		Details:  `{{ "Example:" | faint }} {{ .Example }}`,
	}

	prompt := promptui.Select{
		Label:     label,
		Items:     items,
		Templates: templates,
	}

	idx, _, err := prompt.Run()
	if err != nil {
		return ""
	}

	return items[idx].Value
}
