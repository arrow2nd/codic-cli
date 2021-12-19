package list

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// Item リストアイテム
type Item struct {
	Name    string
	Value   string
	Example string
}

// List リスト
type List []Item

// SelectList リスト内から1つ選択
func (l *List) SelectList(label string) string {
	templates := &promptui.SelectTemplates{
		Label:    "{{ .Name }}?",
		Active:   `{{ ">" | cyan }} {{ .Name | cyan }}`,
		Inactive: "  {{ .Name }}",
		Selected: `{{ "` + label + `:" | faint }} {{ .Name }}`,
		Details:  `{{ "Example:" | faint }} {{ .Example }}`,
	}

	prompt := promptui.Select{
		Label:     label,
		Items:     *l,
		Templates: templates,
	}

	idx, _, err := prompt.Run()
	if err != nil {
		return ""
	}

	return (*l)[idx].Value
}

// GetValue 名前から値を取得
func (l *List) GetValue(name string) (string, error) {
	for _, item := range *l {
		if item.Name == name {
			return item.Value, nil
		}
	}

	return "", fmt.Errorf("does not exist")
}

// GetListItemsString リスト内の全てのアイテム名を文字列で取得
func (l *List) GetListItemsString() string {
	result := ""

	for _, item := range *l {
		result += item.Name + ", "
	}

	return result[:len(result)-2]
}
