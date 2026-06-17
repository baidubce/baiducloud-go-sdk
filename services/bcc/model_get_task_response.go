package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetTaskResponse struct {
	bce.BaseResponse
	Tasks []*TaskDetail `json:"tasks,omitempty"`
}
