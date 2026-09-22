package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ManuallyModifyBandwidthResponse struct {
	bce.BaseResponse
	OrderId *string `json:"orderId,omitempty"`
}
