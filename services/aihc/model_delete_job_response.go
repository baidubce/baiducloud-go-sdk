package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DeleteJobResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
	JobId     *string `json:"jobId,omitempty"`
	JobName   *string `json:"jobName,omitempty"`
}
