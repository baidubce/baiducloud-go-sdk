package ocr

type DocAnalysisLocation struct {
	Left   *int32 `json:"left,omitempty"`
	Top    *int32 `json:"top,omitempty"`
	Width  *int32 `json:"width,omitempty"`
	Height *int32 `json:"height,omitempty"`
}
