package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type EnterRescueModeResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
}
