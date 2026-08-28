package ocr

type DocAnalysisFormulaResult struct {
	FormLocation *FormLocation `json:"form_location,omitempty"`
	FormWords    *string       `json:"form_words,omitempty"`
}
