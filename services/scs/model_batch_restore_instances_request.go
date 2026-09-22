package scs

type BatchRestoreInstancesRequest struct {
	InstanceIds []*string `json:"instanceIds,omitempty"`
}
