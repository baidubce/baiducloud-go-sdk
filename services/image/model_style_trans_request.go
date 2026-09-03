package image

type StyleTransRequest struct {
	Image  *string `json:"image,omitempty"`
	Url    *string `json:"url,omitempty"`
	Option *string `json:"option,omitempty"`
}
