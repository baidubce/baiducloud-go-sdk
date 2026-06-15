package bcc

type BatchStopInstanceRequest struct {
	InstanceIds      []*string `json:"instanceIds,omitempty"`
	ForceStop        *bool     `json:"forceStop,omitempty"`
	StopWithNoCharge *bool     `json:"stopWithNoCharge,omitempty"`
}
