package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ResizeInstanceResponse struct {
	bce.BaseResponse
	OrderId *string `json:"orderId,omitempty"`
}
