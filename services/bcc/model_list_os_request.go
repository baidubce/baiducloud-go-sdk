package bcc

type ListOsRequest struct {
	InstanceIds []*string `json:"instanceIds,omitempty"`
}
