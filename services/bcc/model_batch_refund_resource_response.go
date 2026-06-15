package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BatchRefundResourceResponse struct {
	bce.BaseResponse
	FailedInstanceIds []*string `json:"failedInstanceIds,omitempty"`
}
