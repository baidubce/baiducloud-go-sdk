package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type PurchaseReservedVolumeClusterResponse struct {
	bce.BaseResponse
	OrderId *string `json:"orderId,omitempty"`
}
