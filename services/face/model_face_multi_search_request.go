package face

type FaceMultiSearchRequest struct {
	Image           *string `json:"image,omitempty"`
	ImageType       *string `json:"image_type,omitempty"`
	GroupIdList     *string `json:"group_id_list,omitempty"`
	MaxFaceNum      *int32  `json:"max_face_num,omitempty"`
	MatchThreshold  *int32  `json:"match_threshold,omitempty"`
	QualityControl  *string `json:"quality_control,omitempty"`
	LivenessControl *string `json:"liveness_control,omitempty"`
	SpoofingControl *string `json:"spoofing_control,omitempty"`
	MaxUserNum      *int32  `json:"max_user_num,omitempty"`
}
