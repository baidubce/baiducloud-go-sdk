package face

type FaceSearchRequest struct {
	Image           *string `json:"image,omitempty"`
	ImageType       *string `json:"image_type,omitempty"`
	GroupIdList     *string `json:"group_id_list,omitempty"`
	QualityControl  *string `json:"quality_control,omitempty"`
	LivenessControl *string `json:"liveness_control,omitempty"`
	SpoofingControl *string `json:"spoofing_control,omitempty"`
	UserId          *string `json:"user_id,omitempty"`
	MaxUserNum      *int32  `json:"max_user_num,omitempty"`
	FaceSortType    *int32  `json:"face_sort_type,omitempty"`
	MatchThreshold  *int32  `json:"match_threshold,omitempty"`
}
