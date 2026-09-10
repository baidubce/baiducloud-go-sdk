package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type StepsToObtainNodeEventsV2Response struct {
	bce.BaseResponse
	Status    *string `json:"status,omitempty"`
	Steps     []*Step `json:"steps,omitempty"`
	RequestID *string `json:"requestID,omitempty"`
}
