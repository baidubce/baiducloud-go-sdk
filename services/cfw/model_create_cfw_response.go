package cfw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateCfwResponse struct {
	bce.BaseResponse
	CfwId *string `json:"cfwId,omitempty"`
}
