package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateJobResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
	JobId     *string `json:"jobId,omitempty"`
	JobName   *string `json:"jobName,omitempty"`
}
