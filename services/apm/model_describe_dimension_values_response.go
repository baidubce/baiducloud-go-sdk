package apm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeDimensionValuesResponse struct {
	bce.BaseResponse
	Success *bool     `json:"success,omitempty"`
	Code    *string   `json:"code,omitempty"`
	Message *string   `json:"message,omitempty"`
	Values  []*string `json:"values,omitempty"`
}
