package image

type PicturebookImageDeleteRequest struct {
	Image    *string `json:"image,omitempty"`
	Url      *string `json:"url,omitempty"`
	ContSign *string `json:"cont_sign,omitempty"`
}
