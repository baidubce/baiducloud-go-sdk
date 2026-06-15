package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetInstanceVncResponse struct {
	bce.BaseResponse
	VncUrl *string `json:"vncUrl,omitempty"`
}
