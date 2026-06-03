package pfs

type UpdPerL2BktLnkInfoRequest struct {
	Action                  *string `json:"-"`
	InstanceId              *string `json:"instanceId,omitempty"`
	BucketLinkId            *string `json:"bucketLinkId,omitempty"`
	NewCron                 *string `json:"newCron,omitempty"`
	NewBucketLinkName       *string `json:"newBucketLinkName,omitempty"`
	NewConflictPolicy       *int32  `json:"newConflictPolicy,omitempty"`
	NewThroughputLimitBytes *int32  `json:"newThroughputLimitBytes,omitempty"`
	NewScope                *int32  `json:"newScope,omitempty"`
}
