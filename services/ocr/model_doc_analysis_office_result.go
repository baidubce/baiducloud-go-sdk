package ocr

type DocAnalysisOfficeResult struct {
	WordsType *string                  `json:"words_type,omitempty"`
	Words     *DocAnalysisOfficeWords  `json:"words,omitempty"`
	Chars     []*DocAnalysisOfficeChar `json:"chars,omitempty"`
}
