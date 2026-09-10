package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryTheConfigurationOfAutoscalerV2Response struct {
	bce.BaseResponse
	Autoscaler *Autoscaler `json:"autoscaler,omitempty"`
	RequestID  *string     `json:"requestID,omitempty"`
}
