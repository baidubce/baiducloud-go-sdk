package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ChangeToPrepaidResponse struct {
	bce.BaseResponse
	OrderId *string `json:"orderId,omitempty"`
}
