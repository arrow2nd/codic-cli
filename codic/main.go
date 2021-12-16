package codic

// Codic codic
type Codic struct {
	apiKey string
}

// NewClient 作成
func NewClient(apiKey string) *Codic {
	return &Codic{
		apiKey: apiKey,
	}
}
