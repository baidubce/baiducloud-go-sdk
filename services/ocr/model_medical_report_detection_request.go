package ocr

type MedicalReportDetectionRequest struct {
	Image       *string `json:"image,omitempty"`
	Url         *string `json:"url,omitempty"`
	Location    *bool   `json:"location,omitempty"`
	Probability *bool   `json:"probability,omitempty"`
}
