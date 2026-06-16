package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdateDeploySetRelationResponse struct {
	bce.BaseResponse
	TaskId *string `json:"taskId,omitempty"`
}
