package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetSnapshotResponse struct {
	bce.BaseResponse
	Snapshot *SnapshotModel `json:"snapshot,omitempty"`
}
