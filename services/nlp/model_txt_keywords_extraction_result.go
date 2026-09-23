package nlp

type TxtKeywordsExtractionResult struct {
	Score *float32 `json:"score,omitempty"`
	Word  *string  `json:"word,omitempty"`
}
