package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdateAutoscalerConfigurationV2Response struct {
	bce.BaseResponse
	RequestID *string `json:"requestID,omitempty"`
}
