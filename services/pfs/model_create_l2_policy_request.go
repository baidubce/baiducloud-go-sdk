package pfs

type CreateL2PolicyRequest struct {
	Action       *string `json:"-"`
	InstanceId   *string `json:"instanceId,omitempty"`
	PolicyName   *string `json:"policyName,omitempty"`
	Path         *string `json:"path,omitempty"`
	ExpiredTime  *int32  `json:"expiredTime,omitempty"`
	PfsType      *int32  `json:"type,omitempty"`
	ExecuteTime  *int32  `json:"executeTime,omitempty"`
	BucketName   *string `json:"bucketName,omitempty"`
	BucketPrefix *string `json:"bucketPrefix,omitempty"`
}
