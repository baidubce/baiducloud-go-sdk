package face

type FaceLandmarkRequest struct {
	Image      *string `json:"image,omitempty"`
	ImageType  *string `json:"image_type,omitempty"`
	MaxFaceNum *int32  `json:"max_face_num,omitempty"`
	FaceField  *string `json:"face_field,omitempty"`
}
