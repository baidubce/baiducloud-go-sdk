package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribePodEventsResponse struct {
	bce.BaseResponse
	RequestId *string  `json:"requestId,omitempty"`
	Events    []*Event `json:"events,omitempty"`
	Total     *int32   `json:"total,omitempty"`
}
