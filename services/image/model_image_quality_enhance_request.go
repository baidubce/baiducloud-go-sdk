package image

type ImageQualityEnhanceRequest struct {
	Image *string `json:"image,omitempty"`
	Url   *string `json:"url,omitempty"`
}
