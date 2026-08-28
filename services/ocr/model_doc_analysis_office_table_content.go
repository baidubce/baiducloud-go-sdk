package ocr

type DocAnalysisOfficeTableContent struct {
	PolyLocation []*DocAnalysisOfficePoint `json:"poly_location,omitempty"`
	Word         *string                   `json:"word,omitempty"`
}
