package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeOrderResponse struct {
	bce.BaseResponse
	OrderInfo *OrderInfo `json:"orderInfo,omitempty"`
}
