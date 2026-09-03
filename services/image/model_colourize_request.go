package image

type ColourizeRequest struct {
	Image *string `json:"image,omitempty"`
	Url   *string `json:"url,omitempty"`
}
