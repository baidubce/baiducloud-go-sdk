package ocr

type BusinessLicenseRequest struct {
	Image          *string `json:"image,omitempty"`
	Url            *string `json:"url,omitempty"`
	Accuracy       *string `json:"accuracy,omitempty"`
	RiskWarn       *bool   `json:"risk_warn,omitempty"`
	DetectQuality  *bool   `json:"detect_quality,omitempty"`
	FullwidthShift *bool   `json:"fullwidth_shift,omitempty"`
}
