package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ViewTaskDetailsV2Response struct {
	bce.BaseResponse
	Task      *Task   `json:"task,omitempty"`
	RequestID *string `json:"requestID,omitempty"`
}
