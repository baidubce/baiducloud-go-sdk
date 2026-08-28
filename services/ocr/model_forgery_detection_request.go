package ocr

type ForgeryDetectionRequest struct {
	Image               *string  `json:"image,omitempty"`
	Url                 *string  `json:"url,omitempty"`
	DetectProportion    *bool    `json:"detect_proportion,omitempty"`
	DetectThreshold     *float64 `json:"detect_threshold,omitempty"`
	ReturnHeatmap       *bool    `json:"return_heatmap,omitempty"`
	RestrictProbability *float64 `json:"restrict_probability,omitempty"`
}
