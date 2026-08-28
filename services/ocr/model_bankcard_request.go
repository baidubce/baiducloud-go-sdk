package ocr

type BankcardRequest struct {
	Image         *string `json:"image,omitempty"`
	Url           *string `json:"url,omitempty"`
	Location      *bool   `json:"location,omitempty"`
	DetectQuality *bool   `json:"detect_quality,omitempty"`
}
