package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ModifyJobResponse struct {
	bce.BaseResponse
	JobId     *string `json:"jobId,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
}
