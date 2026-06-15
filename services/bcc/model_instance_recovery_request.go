package bcc

type InstanceRecoveryRequest struct {
	InstanceIds []*InstanceIdItem `json:"instanceIds,omitempty"`
}
