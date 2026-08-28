package ocr

type FormulaResult struct {
	FormLocation []*DocAnalysisOfficePoint `json:"form_location,omitempty"`
	FormWords    *string                   `json:"form_words,omitempty"`
}
