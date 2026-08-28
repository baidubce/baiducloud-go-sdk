package ocr

type DocAnalysisWordsResult struct {
	Location *DocAnalysisLocation `json:"location,omitempty"`
	Words    *string              `json:"words,omitempty"`
	OcrType  *string              `json:"type,omitempty"`
	Chars    []*DocAnalysisChar   `json:"chars,omitempty"`
}
