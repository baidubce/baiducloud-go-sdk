package face

type FaceEditAttrRequest struct {
	Image          *string `json:"image,omitempty"`
	ImageType      *string `json:"image_type,omitempty"`
	ActionType     *string `json:"action_type,omitempty"`
	Target         *int32  `json:"target,omitempty"`
	QualityControl *string `json:"quality_control,omitempty"`
	FaceLocation   *string `json:"face_location,omitempty"`
}
