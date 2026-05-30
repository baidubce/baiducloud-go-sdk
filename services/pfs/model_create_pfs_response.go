package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreatePfsResponse struct {
	bce.BaseResponse
	InstanceId *string `json:"instanceId,omitempty"`
	OrderId    *string `json:"orderId,omitempty"`
}
