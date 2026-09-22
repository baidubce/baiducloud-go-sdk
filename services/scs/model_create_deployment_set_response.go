package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateDeploymentSetResponse struct {
	bce.BaseResponse
	DeploySetId *string `json:"deploySetId,omitempty"`
}
