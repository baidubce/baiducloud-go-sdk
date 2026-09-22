package face

type UserDeleteRequest struct {
	GroupId *string `json:"group_id,omitempty"`
	UserId  *string `json:"user_id,omitempty"`
}
