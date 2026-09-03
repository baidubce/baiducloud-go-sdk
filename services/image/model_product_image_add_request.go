package image

type ProductImageAddRequest struct {
	Image    *string `json:"image,omitempty"`
	Url      *string `json:"url,omitempty"`
	Brief    *string `json:"brief,omitempty"`
	ClassId1 *int32  `json:"class_id1,omitempty"`
	ClassId2 *int32  `json:"class_id2,omitempty"`
}
