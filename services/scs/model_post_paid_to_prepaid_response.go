package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type PostPaidToPrepaidResponse struct {
	bce.BaseResponse
	OrderId *string `json:"orderId,omitempty"`
}
