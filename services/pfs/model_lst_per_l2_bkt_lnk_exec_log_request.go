package pfs

type LstPerL2BktLnkExecLogRequest struct {
	Action       *string `json:"-"`
	InstanceId   *string `json:"instanceId,omitempty"`
	BucketLinkId *string `json:"bucketLinkId,omitempty"`
	StartTime    *int32  `json:"startTime,omitempty"`
	EndTime      *int32  `json:"endTime,omitempty"`
}
