package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetTaskListV2Response struct {
	bce.BaseResponse
	Page      *ListTaskPage `json:"page,omitempty"`
	RequestID *string       `json:"requestID,omitempty"`
}
