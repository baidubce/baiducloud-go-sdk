package image

type LogoRequest struct {
	Image     *string `json:"image,omitempty"`
	Url       *string `json:"url,omitempty"`
	CustomLib *bool   `json:"custom_lib,omitempty"`
}
