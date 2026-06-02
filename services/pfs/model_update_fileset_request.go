package pfs

type UpdateFilesetRequest struct {
	Action           *string `json:"-"`
	InstanceId       *string `json:"instanceId,omitempty"`
	FilesetId        *string `json:"filesetId,omitempty"`
	FilesetName      *string `json:"filesetName,omitempty"`
	BlockQuota       *int32  `json:"blockQuota,omitempty"`
	FilesQuota       *int64  `json:"filesQuota,omitempty"`
	QpsLimit         *int32  `json:"qpsLimit,omitempty"`
	BandwidthLimitMb *int32  `json:"bandwidthLimitMb,omitempty"`
}
