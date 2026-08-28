package ocr

type ForgeryDetectionResult struct {
	DetectionResult    *string             `json:"detection_result,omitempty"`
	TamperedProportion *float32            `json:"tampered_proportion,omitempty"`
	TamperedLocation   []*TamperedLocation `json:"tampered_location,omitempty"`
	Heatmap            *string             `json:"heatmap,omitempty"`
}
