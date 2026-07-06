package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeJobsResponse struct {
	bce.BaseResponse
	RequestId  *string    `json:"requestId,omitempty"`
	TotalCount *int32     `json:"totalCount,omitempty"`
	Jobs       []*JobItem `json:"jobs,omitempty"`
}
