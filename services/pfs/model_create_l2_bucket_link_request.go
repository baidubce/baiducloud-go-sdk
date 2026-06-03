package pfs

type CreateL2BucketLinkRequest struct {
	Action               *string `json:"-"`
	InstanceId           *string `json:"instanceId,omitempty"`
	ConflictPolicy       *string `json:"conflictPolicy,omitempty"`
	BucketName           *string `json:"bucketName,omitempty"`
	BucketPrefix         *string `json:"bucketPrefix,omitempty"`
	ThroughputLimitBytes *string `json:"throughputLimitBytes,omitempty"`
	ReportObjectName     *string `json:"reportObjectName,omitempty"`
	BucketLinkName       *string `json:"bucketLinkName,omitempty"`
	TransferType         *int32  `json:"transferType,omitempty"`
	PfsPath              *string `json:"pfsPath,omitempty"`
	Cron                 *string `json:"cron,omitempty"`
	BucketBelongUserId   *string `json:"bucketBelongUserId,omitempty"`
	LccId                *string `json:"lccId,omitempty"`
	Scope                *int32  `json:"scope,omitempty"`
}
