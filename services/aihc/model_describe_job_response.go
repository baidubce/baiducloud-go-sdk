package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeJobResponse struct {
	bce.BaseResponse
	RequestId *string  `json:"requestId,omitempty"`
	Job       *JobItem `json:"job,omitempty"`
}
