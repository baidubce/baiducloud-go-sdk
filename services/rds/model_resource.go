package rds

type Resource struct {
	InstanceId *string `json:"instanceId,omitempty"`
	Tags       []*Tag  `json:"tags,omitempty"`
}
