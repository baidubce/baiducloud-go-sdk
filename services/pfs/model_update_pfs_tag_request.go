package pfs

type UpdatePFSTagRequest struct {
	InstanceId []*string `json:"instanceId,omitempty"`
	Tags       []*Tag    `json:"tags,omitempty"`
}
