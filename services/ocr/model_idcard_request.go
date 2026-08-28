package ocr

type IdcardRequest struct {
	IdCardSide       *string `json:"id_card_side,omitempty"`
	Image            *string `json:"image,omitempty"`
	Url              *string `json:"url,omitempty"`
	DetectPs         *bool   `json:"detect_ps,omitempty"`
	DetectRisk       *bool   `json:"detect_risk,omitempty"`
	DetectQuality    *bool   `json:"detect_quality,omitempty"`
	DetectPhoto      *bool   `json:"detect_photo,omitempty"`
	DetectCard       *bool   `json:"detect_card,omitempty"`
	DetectDirection  *bool   `json:"detect_direction,omitempty"`
	DetectScreenshot *bool   `json:"detect_screenshot,omitempty"`
}
