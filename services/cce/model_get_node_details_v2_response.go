package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetNodeDetailsV2Response struct {
	bce.BaseResponse
	Instance  *Instance `json:"instance,omitempty"`
	RequestID *string   `json:"requestID,omitempty"`
}
