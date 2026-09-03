package image

type LogoDeleteRequest struct {
	Image    *string `json:"image,omitempty"`
	ContSign *string `json:"cont_sign,omitempty"`
}
