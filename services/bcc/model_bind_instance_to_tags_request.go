package bcc

type BindInstanceToTagsRequest struct {
	InstanceId               *string     `json:"-"`
	ChangeTags               []*TagModel `json:"changeTags,omitempty"`
	AttachRelatedResourceTag *bool       `json:"attachRelatedResourceTag,omitempty"`
}
