package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ReleaseInstanceResponse struct {
	bce.BaseResponse
	Success                  *bool     `json:"success,omitempty"`
	Message                  *string   `json:"message,omitempty"`
	ReleaseFailedInstanceIds []*string `json:"releaseFailedInstanceIds,omitempty"`
}
