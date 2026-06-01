package apm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeLLMSessionsStatisticsResponse struct {
	bce.BaseResponse
	Success *bool   `json:"success,omitempty"`
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Count   *int32  `json:"count,omitempty"`
}
