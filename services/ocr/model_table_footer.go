package ocr

type TableFooter struct {
	Location []*DocAnalysisOfficePoint `json:"location,omitempty"`
	Words    *string                   `json:"words,omitempty"`
}
