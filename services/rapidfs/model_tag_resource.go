package rapidfs

type TagResource struct {
	InstanceId *string `json:"instanceId,omitempty"`
	Tags       []*Tag  `json:"tags,omitempty"`
}
