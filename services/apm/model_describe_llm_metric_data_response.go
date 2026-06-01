package apm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeLLMMetricDataResponse struct {
	bce.BaseResponse
	Success    *bool            `json:"success,omitempty"`
	Code       *string          `json:"code,omitempty"`
	Message    *string          `json:"message,omitempty"`
	Timeseries []*Timeseries    `json:"timeseries,omitempty"`
	Aggregate  *AggregateResult `json:"aggregate,omitempty"`
}
