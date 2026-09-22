package face

type FaceDetectRequest struct {
	Image            *string `json:"image,omitempty"`
	ImageType        *string `json:"image_type,omitempty"`
	FaceField        *string `json:"face_field,omitempty"`
	MaxFaceNum       *int32  `json:"max_face_num,omitempty"`
	FaceType         *string `json:"face_type,omitempty"`
	LivenessControl  *string `json:"liveness_control,omitempty"`
	FaceSortType     *int32  `json:"face_sort_type,omitempty"`
	DisplayCorpImage *int32  `json:"display_corp_image,omitempty"`
}
