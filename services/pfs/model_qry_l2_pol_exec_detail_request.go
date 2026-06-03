package pfs

type QryL2PolExecDetailRequest struct {
	Action     *string `json:"-"`
	InstanceId *string `json:"instanceId,omitempty"`
	PolicyId   *string `json:"policyId,omitempty"`
	JobId      *string `json:"jobId,omitempty"`
}
