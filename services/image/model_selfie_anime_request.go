package image

type SelfieAnimeRequest struct {
	Image     *string `json:"image,omitempty"`
	Url       *string `json:"url,omitempty"`
	ImageType *string `json:"type,omitempty"`
	MaskId    *string `json:"mask_id,omitempty"`
}
