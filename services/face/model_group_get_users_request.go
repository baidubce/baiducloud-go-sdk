package face

type GroupGetUsersRequest struct {
	GroupId *string `json:"group_id,omitempty"`
	Start   *int32  `json:"start,omitempty"`
	Length  *int32  `json:"length,omitempty"`
}
