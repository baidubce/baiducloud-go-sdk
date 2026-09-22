package scs

type UnbindTagsRequest struct {
	InstanceId *string     `json:"-"`
	ChangeTags []*TagModel `json:"changeTags,omitempty"`
}
