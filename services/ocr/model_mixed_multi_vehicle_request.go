package ocr

type MixedMultiVehicleRequest struct {
	Image           *string `json:"image,omitempty"`
	Url             *string `json:"url,omitempty"`
	DetectDirection *bool   `json:"detect_direction,omitempty"`
	Unified         *bool   `json:"unified,omitempty"`
}
