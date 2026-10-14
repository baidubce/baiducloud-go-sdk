package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryTaskListResponse struct {
	bce.BaseResponse
	Tasks []*Task `json:"tasks,omitempty"`
	Count *int32  `json:"count,omitempty"`
}
