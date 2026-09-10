package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateExpansionNodeGroupTaskV2Response struct {
	bce.BaseResponse
	TaskID    *string `json:"taskID,omitempty"`
	RequestID *string `json:"requestID,omitempty"`
}
