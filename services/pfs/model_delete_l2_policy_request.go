package pfs

type DeleteL2PolicyRequest struct {
	Action     *string `json:"-"`
	InstanceId *string `json:"instanceId,omitempty"`
	PolicyId   *string `json:"policyId,omitempty"`
}
