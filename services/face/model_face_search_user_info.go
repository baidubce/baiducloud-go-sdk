package face

type FaceSearchUserInfo struct {
	Score    *float32 `json:"score,omitempty"`
	GroupId  *string  `json:"group_id,omitempty"`
	UserId   *string  `json:"user_id,omitempty"`
	UserInfo *string  `json:"user_info,omitempty"`
}
