package pfs

type DescL2PolicyRequest struct {
	Action     *string `json:"-"`
	InstanceId *string `json:"instanceId,omitempty"`
	PolicyId   *string `json:"policyId,omitempty"`
}
