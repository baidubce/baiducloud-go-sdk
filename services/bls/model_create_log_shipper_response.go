package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateLogShipperResponse struct {
	bce.BaseResponse
	Success *bool   `json:"success,omitempty"`
	Code    *string `json:"code,omitempty"`
}
