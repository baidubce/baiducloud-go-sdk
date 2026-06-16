package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateDeploySetResponse struct {
	bce.BaseResponse
	DeploySetIds []*string `json:"deploySetIds,omitempty"`
}
