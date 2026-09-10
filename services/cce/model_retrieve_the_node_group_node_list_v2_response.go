package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type RetrieveTheNodeGroupNodeListV2Response struct {
	bce.BaseResponse
	Page      *ListInstancesByInstanceGroupIDPage `json:"page,omitempty"`
	RequestID *string                             `json:"requestID,omitempty"`
}
