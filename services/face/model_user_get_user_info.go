package face

type UserGetUserInfo struct {
	UserInfo *string `json:"user_info,omitempty"`
	GroupId  *string `json:"group_id,omitempty"`
}
