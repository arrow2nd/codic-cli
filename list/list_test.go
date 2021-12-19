package list

import (
	"testing"
)

func TestList_GetValue(t *testing.T) {
	straylight := &List{
		{
			Name:    "asahi",
			Value:   "芹沢あさひ",
			Example: "Straylight",
		},
		{
			Name:    "fuyuko",
			Value:   "黛冬優子",
			Example: "Straylight",
		},
		{
			Name:    "mei",
			Value:   "和泉愛依",
			Example: "Straylight",
		},
	}

	tests := []struct {
		name     string
		l        *List
		itemName string
		want     string
		wantErr  bool
	}{
		{
			name:     "値を取得できるか",
			l:        straylight,
			itemName: "asahi",
			want:     "芹沢あさひ",
			wantErr:  false,
		},
		{
			name:     "存在しない名前を指定したときエラーを返すか",
			l:        straylight,
			itemName: "mano",
			want:     "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.GetValue(tt.itemName)
			if (err != nil) != tt.wantErr {
				t.Errorf("List.GetValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("List.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_GetListItemsString(t *testing.T) {
	tests := []struct {
		name string
		l    *List
		want string
	}{
		{
			name: "取得できるか (SHHis)",
			l: &List{
				{
					Name:    "nichika",
					Value:   "七草にちか",
					Example: "SHHis",
				},
				{
					Name:    "mikoto",
					Value:   "緋田美琴",
					Example: "SHHis",
				},
			},
			want: "nichika, mikoto",
		},
		{
			name: "取得できるか (noctchill)",
			l: &List{
				{
					Name:    "toru",
					Value:   "浅倉透",
					Example: "noctchill",
				},
				{
					Name:    "madoka",
					Value:   "樋口円香",
					Example: "noctchill",
				},
				{
					Name:    "koito",
					Value:   "福丸小糸",
					Example: "noctchill",
				},
				{
					Name:    "hinana",
					Value:   "市川雛菜",
					Example: "noctchill",
				},
			},
			want: "toru, madoka, koito, hinana",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.GetListItemsString(); got != tt.want {
				t.Errorf("List.GetListItemsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
