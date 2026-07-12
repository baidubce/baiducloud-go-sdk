package oos

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdateTemplateV2Response struct {
	bce.BaseResponse
	Success *bool   `json:"success,omitempty"`
	Msg     *string `json:"msg,omitempty"`
}
