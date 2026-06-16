package cfw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateStatelessCfwResponse struct {
	bce.BaseResponse
	CfwId *string `json:"cfwId,omitempty"`
}
