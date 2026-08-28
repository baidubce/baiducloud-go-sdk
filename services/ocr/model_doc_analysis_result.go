package ocr

type DocAnalysisResult struct {
	WordsType *string            `json:"words_type,omitempty"`
	Words     *DocAnalysisWords  `json:"words,omitempty"`
	Chars     []*DocAnalysisChar `json:"chars,omitempty"`
}
