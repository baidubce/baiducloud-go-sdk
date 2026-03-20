package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BillingChangePreToPostBlbResponse struct {
	bce.BaseResponse
	OrderId *string `json:"orderId,omitempty"`
}
