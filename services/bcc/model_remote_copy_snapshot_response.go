package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type RemoteCopySnapshotResponse struct {
	bce.BaseResponse
	Result []*RemoteCopySnapshot `json:"result,omitempty"`
}
