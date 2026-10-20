package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type RenewalExampleResponse struct {
	bce.BaseResponse
	OrderId *string `json:"orderId,omitempty"`
}
