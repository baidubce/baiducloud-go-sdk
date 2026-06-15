package bcc

type StopInstanceRequest struct {
	InstanceId       *string `json:"-"`
	ForceStop        *bool   `json:"forceStop,omitempty"`
	StopWithNoCharge *bool   `json:"stopWithNoCharge,omitempty"`
}
