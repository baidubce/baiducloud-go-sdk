package ccr

type ExecuteImageMigrationRequest struct {
	InstanceId *string `json:"-"`
	PolicyId   *int32  `json:"policyId,omitempty"`
}
