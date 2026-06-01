package apm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeEnvResponse struct {
	bce.BaseResponse
	Success *bool     `json:"success,omitempty"`
	Code    *string   `json:"code,omitempty"`
	Message *string   `json:"message,omitempty"`
	Env     []*string `json:"env,omitempty"`
}
