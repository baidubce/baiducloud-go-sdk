package oos

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateExecutionV2Response struct {
	bce.BaseResponse
	Success *bool        `json:"success,omitempty"`
	Msg     *string      `json:"msg,omitempty"`
	Code    *int32       `json:"code,omitempty"`
	Result  *interface{} `json:"result,omitempty"`
}
