package ocr

type FacadeWordsResult struct {
	Words *string  `json:"words,omitempty"`
	Score *float32 `json:"score,omitempty"`
	Brief *string  `json:"brief,omitempty"`
}
