package ocr

type MultiIdcardRequest struct {
	Image            *string `json:"image,omitempty"`
	Url              *string `json:"url,omitempty"`
	DetectRisk       *bool   `json:"detect_risk,omitempty"`
	DetectQuality    *bool   `json:"detect_quality,omitempty"`
	DetectPhoto      *bool   `json:"detect_photo,omitempty"`
	DetectCard       *bool   `json:"detect_card,omitempty"`
	DetectScreenshot *bool   `json:"detect_screenshot,omitempty"`
}
