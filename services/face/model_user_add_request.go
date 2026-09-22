package face

type UserAddRequest struct {
	Image           *string `json:"image,omitempty"`
	ImageType       *string `json:"image_type,omitempty"`
	GroupId         *string `json:"group_id,omitempty"`
	UserId          *string `json:"user_id,omitempty"`
	UserInfo        *string `json:"user_info,omitempty"`
	QualityControl  *string `json:"quality_control,omitempty"`
	LivenessControl *string `json:"liveness_control,omitempty"`
	SpoofingControl *string `json:"spoofing_control,omitempty"`
	ActionType      *string `json:"action_type,omitempty"`
	FaceSortType    *int32  `json:"face_sort_type,omitempty"`
}
