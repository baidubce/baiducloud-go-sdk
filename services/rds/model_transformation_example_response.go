package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type TransformationExampleResponse struct {
	bce.BaseResponse
	OrderId *string `json:"orderId,omitempty"`
}
