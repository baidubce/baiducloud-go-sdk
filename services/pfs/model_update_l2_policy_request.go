package pfs

type UpdateL2PolicyRequest struct {
	Action        *string `json:"-"`
	InstanceId    *string `json:"instanceId,omitempty"`
	PolicyId      *string `json:"policyId,omitempty"`
	NewPolicyName *string `json:"newPolicyName,omitempty"`
	ExpiredTime   *int32  `json:"expiredTime,omitempty"`
	ExecuteTime   *int32  `json:"executeTime,omitempty"`
	BucketName    *string `json:"bucketName,omitempty"`
	BucketPrefix  *string `json:"bucketPrefix,omitempty"`
}
