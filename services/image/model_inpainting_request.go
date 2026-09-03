package image

type InpaintingRequest struct {
	Image     *string `json:"image,omitempty"`
	Url       *string `json:"url,omitempty"`
	Rectangle *string `json:"rectangle,omitempty"`
}
