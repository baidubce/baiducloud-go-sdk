package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeJobLogsResponse struct {
	bce.BaseResponse
	RequestId  *string   `json:"requestId,omitempty"`
	JobId      *string   `json:"jobId,omitempty"`
	PodName    *string   `json:"podName,omitempty"`
	Logs       []*string `json:"logs,omitempty"`
	NextMarker *string   `json:"nextMarker,omitempty"`
}
