package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BatchChangeToPostpaidResponse struct {
	bce.BaseResponse
	OrderId *string `json:"orderId,omitempty"`
}
