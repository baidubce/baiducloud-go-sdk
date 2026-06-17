package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateTaskResponse struct {
	bce.BaseResponse
	TaskId *string `json:"taskId,omitempty"`
}
