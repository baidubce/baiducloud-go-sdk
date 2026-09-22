package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type RenewInstanceResponse struct {
	bce.BaseResponse
	OrderId *string `json:"orderId,omitempty"`
}
