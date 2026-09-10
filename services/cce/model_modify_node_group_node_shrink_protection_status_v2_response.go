package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ModifyNodeGroupNodeShrinkProtectionStatusV2Response struct {
	bce.BaseResponse
	FailedInstances []*interface{} `json:"failedInstances,omitempty"`
	RequestID       *string        `json:"requestID,omitempty"`
}
