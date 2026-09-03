package image

type ContrastEnhanceRequest struct {
	Image *string `json:"image,omitempty"`
	Url   *string `json:"url,omitempty"`
}
