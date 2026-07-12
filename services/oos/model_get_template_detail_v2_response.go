package oos

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetTemplateDetailV2Response struct {
	bce.BaseResponse
	Success *bool     `json:"success,omitempty"`
	Msg     *string   `json:"msg,omitempty"`
	Code    *int32    `json:"code,omitempty"`
	Result  *Template `json:"result,omitempty"`
}
