package apm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeRetentionLimitResponse struct {
	bce.BaseResponse
	Success                *bool   `json:"success,omitempty"`
	Code                   *string `json:"code,omitempty"`
	Message                *string `json:"message,omitempty"`
	MinTraceRetentionDays  *int32  `json:"minTraceRetentionDays,omitempty"`
	MaxTraceRetentionDays  *int32  `json:"maxTraceRetentionDays,omitempty"`
	MinMetricRetentionDays *int32  `json:"minMetricRetentionDays,omitempty"`
	MaxMetricRetentionDays *int32  `json:"maxMetricRetentionDays,omitempty"`
	MinDorisRetentionDays  *int32  `json:"minDorisRetentionDays,omitempty"`
	MaxDorisRetentionDays  *int32  `json:"maxDorisRetentionDays,omitempty"`
}
