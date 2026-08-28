package ocr

type DocAnalysisOfficeWords struct {
	Word            *string                           `json:"word,omitempty"`
	LineProbability *DocAnalysisOfficeLineProbability `json:"line_probability,omitempty"`
	PolyLocation    []*DocAnalysisOfficePoint         `json:"poly_location,omitempty"`
	WordsLocation   *DocAnaysisOfficeLocation         `json:"words_location,omitempty"`
}
