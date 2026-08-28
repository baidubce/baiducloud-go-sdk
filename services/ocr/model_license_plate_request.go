package ocr

type LicensePlateRequest struct {
	Image          *string `json:"image,omitempty"`
	Url            *string `json:"url,omitempty"`
	MultiDetect    *bool   `json:"multi_detect,omitempty"`
	MultiScale     *bool   `json:"multi_scale,omitempty"`
	DetectComplete *bool   `json:"detect_complete,omitempty"`
	DetectRisk     *bool   `json:"detect_risk,omitempty"`
}
