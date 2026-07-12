package oos

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetExecutionDetailV2Response struct {
	bce.BaseResponse
	Success *bool      `json:"success,omitempty"`
	Msg     *string    `json:"msg,omitempty"`
	Code    *int32     `json:"code,omitempty"`
	Result  *Execution `json:"result,omitempty"`
}
