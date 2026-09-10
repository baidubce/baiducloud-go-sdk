package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateNodeGroupV2Response struct {
	bce.BaseResponse
	InstanceGroupID *string `json:"instanceGroupID,omitempty"`
	RequestID       *string `json:"requestID,omitempty"`
}
