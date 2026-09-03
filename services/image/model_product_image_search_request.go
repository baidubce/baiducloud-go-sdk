package image

type ProductImageSearchRequest struct {
	Image    *string `json:"image,omitempty"`
	Url      *string `json:"url,omitempty"`
	ClassId1 *int32  `json:"class_id1,omitempty"`
	ClassId2 *int32  `json:"class_id2,omitempty"`
	TagLogic *int32  `json:"tag_logic,omitempty"`
	Pn       *int32  `json:"pn,omitempty"`
	Rn       *int32  `json:"rn,omitempty"`
}
