package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeMetricDataResponse struct {
	bce.BaseResponse
	Success    *bool         `json:"success,omitempty"`
	Code       *string       `json:"code,omitempty"`
	Message    *string       `json:"message,omitempty"`
	Timeseries []*Timeseries `json:"timeseries,omitempty"`
}
