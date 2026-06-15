package bcc

type UnbindInstanceFromTagsRequest struct {
	InstanceId               *string     `json:"-"`
	ChangeTags               []*TagModel `json:"changeTags,omitempty"`
	AttachRelatedResourceTag *bool       `json:"attachRelatedResourceTag,omitempty"`
}
