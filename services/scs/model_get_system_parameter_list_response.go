package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetSystemParameterListResponse struct {
	bce.BaseResponse
	Success *bool     `json:"success,omitempty"`
	Result  []*Result `json:"result,omitempty"`
}
