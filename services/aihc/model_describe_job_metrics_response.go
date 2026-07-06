package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeJobMetricsResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
	JobId     *string `json:"jobId,omitempty"`
	Metrics   *Metric `json:"metrics,omitempty"`
}
