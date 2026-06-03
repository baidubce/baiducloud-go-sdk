package ccr

type UpdateInstanceTagsRequest struct {
	InstanceId *string       `json:"-"`
	Tags       []*LogicalTag `json:"tags,omitempty"`
}
