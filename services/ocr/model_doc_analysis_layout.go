package ocr

type DocAnalysisLayout struct {
	Layout         *string             `json:"layout,omitempty"`
	LayoutLocation []*DocAnalysisPoint `json:"layout_location,omitempty"`
	LayoutIdx      []*int32            `json:"layout_idx,omitempty"`
}
