package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeJobEventsResponse struct {
	bce.BaseResponse
	RequestId *string  `json:"requestId,omitempty"`
	Events    []*Event `json:"events,omitempty"`
	Total     *int32   `json:"total,omitempty"`
}
