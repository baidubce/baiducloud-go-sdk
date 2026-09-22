package face

type UserGetRequest struct {
	UserId  *string `json:"user_id,omitempty"`
	GroupId *string `json:"group_id,omitempty"`
}
