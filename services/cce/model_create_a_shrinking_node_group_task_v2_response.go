package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAShrinkingNodeGroupTaskV2Response struct {
	bce.BaseResponse
	TaskID    *string `json:"taskID,omitempty"`
	RequestID *string `json:"requestID,omitempty"`
}
