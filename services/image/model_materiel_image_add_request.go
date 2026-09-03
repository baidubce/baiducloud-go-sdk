package image

type MaterielImageAddRequest struct {
	Image *string `json:"image,omitempty"`
	Url   *string `json:"url,omitempty"`
	Brief *string `json:"brief,omitempty"`
	Tags  *string `json:"tags,omitempty"`
}
