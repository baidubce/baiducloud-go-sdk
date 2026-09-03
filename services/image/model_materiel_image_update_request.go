package image

type MaterielImageUpdateRequest struct {
	Image    *string `json:"image,omitempty"`
	Url      *string `json:"url,omitempty"`
	ContSign *string `json:"cont_sign,omitempty"`
	Brief    *string `json:"brief,omitempty"`
	Tags     *string `json:"tags,omitempty"`
}
