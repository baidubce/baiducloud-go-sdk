package ocr

type TableHeader struct {
	Location []*DocAnalysisOfficePoint `json:"location,omitempty"`
	Words    *string                   `json:"words,omitempty"`
}
