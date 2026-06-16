package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateSnapshotShareResponse struct {
	bce.BaseResponse
	SnapshotId *string `json:"snapshotId,omitempty"`
}
