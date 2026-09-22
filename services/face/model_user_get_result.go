package face

type UserGetResult struct {
	UserList []*UserGetUserInfo `json:"user_list,omitempty"`
}
