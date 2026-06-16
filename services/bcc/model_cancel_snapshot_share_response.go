package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CancelSnapshotShareResponse struct {
	bce.BaseResponse
	SourceSnapshotId *string `json:"sourceSnapshotId,omitempty"`
	ShareSnapshotId  *string `json:"shareSnapshotId,omitempty"`
}
