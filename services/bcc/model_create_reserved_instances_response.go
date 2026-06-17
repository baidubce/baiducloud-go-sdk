package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateReservedInstancesResponse struct {
	bce.BaseResponse
	OrderId             *string   `json:"orderId,omitempty"`
	ReservedInstanceIds []*string `json:"reservedInstanceIds,omitempty"`
}
