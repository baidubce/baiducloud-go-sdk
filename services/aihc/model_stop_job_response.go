package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type StopJobResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
	JobId     *string `json:"jobId,omitempty"`
}
