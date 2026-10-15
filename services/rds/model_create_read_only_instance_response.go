package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateReadOnlyInstanceResponse struct {
	bce.BaseResponse
	InstanceIds []*string `json:"instanceIds,omitempty"`
	OrderId     *string   `json:"orderId,omitempty"`
}
