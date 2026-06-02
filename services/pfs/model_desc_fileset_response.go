package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescFilesetResponse struct {
	bce.BaseResponse
	RequestId        *string `json:"requestId,omitempty"`
	FilesetName      *string `json:"filesetName,omitempty"`
	FilesetId        *string `json:"filesetId,omitempty"`
	FilesetPath      *string `json:"filesetPath,omitempty"`
	Status           *int32  `json:"status,omitempty"`
	BlockQuota       *int32  `json:"blockQuota,omitempty"`
	BlockUsage       *int32  `json:"blockUsage,omitempty"`
	FilesQuota       *int32  `json:"filesQuota,omitempty"`
	FilesUsage       *int32  `json:"filesUsage,omitempty"`
	AllocatedInode   *int64  `json:"allocatedInode,omitempty"`
	Ctime            *string `json:"ctime,omitempty"`
	QpsLimit         *int64  `json:"qpsLimit,omitempty"`
	BandwidthLimitMb *int64  `json:"bandwidthLimitMb,omitempty"`
	ParentPath       *bool   `json:"parentPath,omitempty"`
}
