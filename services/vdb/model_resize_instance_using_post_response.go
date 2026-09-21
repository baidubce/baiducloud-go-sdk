package vdb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ResizeInstanceUsingPOSTResponse struct {
	bce.BaseResponse
	OrderId *string `json:"orderId,omitempty"`
}
