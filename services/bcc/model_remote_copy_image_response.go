package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type RemoteCopyImageResponse struct {
	bce.BaseResponse
	Result []*RemoteCopyResult `json:"result,omitempty"`
}
