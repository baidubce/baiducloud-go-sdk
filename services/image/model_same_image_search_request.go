package image

type SameImageSearchRequest struct {
	Image    *string `json:"image,omitempty"`
	Url      *string `json:"url,omitempty"`
	Tags     *string `json:"tags,omitempty"`
	TagLogic *int32  `json:"tag_logic,omitempty"`
	Pn       *int32  `json:"pn,omitempty"`
	Rn       *int32  `json:"rn,omitempty"`
}
