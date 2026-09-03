package image

type ProductImageUpdateRequest struct {
	Image    *string `json:"image,omitempty"`
	Url      *string `json:"url,omitempty"`
	ContSign *string `json:"cont_sign,omitempty"`
	Brief    *string `json:"brief,omitempty"`
	ClassId1 *int32  `json:"class_id1,omitempty"`
	ClassId2 *int32  `json:"class_id2,omitempty"`
}
