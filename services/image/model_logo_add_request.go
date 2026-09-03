package image

type LogoAddRequest struct {
	Image *string `json:"image,omitempty"`
	Url   *string `json:"url,omitempty"`
	Brief *string `json:"brief,omitempty"`
}
