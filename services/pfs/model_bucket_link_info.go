package pfs

type BucketLinkInfo struct {
	BucketlinkId       *string `json:"bucketlinkId,omitempty"`
	BucketlinkName     *string `json:"bucketlinkName,omitempty"`
	InstanceId         *string `json:"instanceId,omitempty"`
	TransferType       *int32  `json:"transferType,omitempty"`
	CreateTime         *string `json:"createTime,omitempty"`
	FinishTime         *string `json:"finishTime,omitempty"`
	ConflictPolicy     *int32  `json:"conflictPolicy,omitempty"`
	ThroughputLimit    *int32  `json:"throughputLimit,omitempty"`
	Status             *int32  `json:"status,omitempty"`
	Src                *string `json:"src,omitempty"`
	Dst                *string `json:"dst,omitempty"`
	Progress           *int32  `json:"progress,omitempty"`
	Report             *string `json:"report,omitempty"`
	Cron               *string `json:"cron,omitempty"`
	BucketBelongUserId *string `json:"bucketBelongUserId,omitempty"`
	LccId              *string `json:"lccId,omitempty"`
	Scope              *int32  `json:"scope,omitempty"`
}
