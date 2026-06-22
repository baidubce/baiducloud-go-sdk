package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeSnapshotsUsageResponse struct {
	bce.BaseResponse
	SnapshotCount       *int32   `json:"snapshotCount,omitempty"`
	AutoSnapshotCount   *int32   `json:"autoSnapshotCount,omitempty"`
	ManualSnapshotCount *int32   `json:"manualSnapshotCount,omitempty"`
	SnapshotCapacity    *float64 `json:"snapshotCapacity,omitempty"`
}
