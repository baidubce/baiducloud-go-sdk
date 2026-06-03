package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateInstanceResponse struct {
	bce.BaseResponse
	OrderId    *string `json:"orderId,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
}
