package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAnInstanceResponse struct {
	bce.BaseResponse
	InstanceIds []*string `json:"instanceIds,omitempty"`
	OrderId     *string   `json:"orderId,omitempty"`
}
