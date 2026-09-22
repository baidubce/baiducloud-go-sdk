package scs

type BindTagsRequest struct {
	InstanceId *string     `json:"-"`
	ChangeTags []*TagModel `json:"changeTags,omitempty"`
}
