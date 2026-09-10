package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ModifyIGAutoScalerResponse struct {
	bce.BaseResponse
	RequestID *string `json:"requestID,omitempty"`
}
