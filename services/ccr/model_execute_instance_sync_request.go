package ccr

type ExecuteInstanceSyncRequest struct {
	InstanceId *string `json:"-"`
	PolicyId   *int32  `json:"policyId,omitempty"`
}
