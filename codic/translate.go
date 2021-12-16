package codic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type base struct {
	Successful bool   `json:"successful"`
	Text       string `json:"text"`
	Result     string `json:"translated_text"`
}

// Candidate 候補
type Candidate struct {
	Text string `json:"text"`
}

// Word 単語
type Word struct {
	base
	Candidates []Candidate `json:"candidates"`
}

// Translate 変換結果
type Translate struct {
	base
	Words []Word `json:"words"`
}

// GetTranslate ネーミング変換結果を取得
func (c *Codic) GetTranslate(params url.Values) ([]Translate, error) {
	req, _ := http.NewRequest("GET", APIBaseURL+"v1/engine/translate.json", nil)
	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	req.URL.RawQuery = params.Encode()

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("access failed (%s)", res.Status)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	results := []Translate{}
	if err := json.Unmarshal(bytes, &results); err != nil {
		return nil, err
	}

	return results, nil
}
