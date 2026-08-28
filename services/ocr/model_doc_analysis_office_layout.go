package ocr

type DocAnalysisOfficeLayout struct {
	Layout         *string                   `json:"layout,omitempty"`
	LayoutProb     *float64                  `json:"layout_prob,omitempty"`
	LayoutLocation []*DocAnalysisOfficePoint `json:"layout_location,omitempty"`
	LayoutIdx      []*int32                  `json:"layout_idx,omitempty"`
}
