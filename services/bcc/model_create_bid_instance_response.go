package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateBidInstanceResponse struct {
	bce.BaseResponse
	OrderId     *string   `json:"orderId,omitempty"`
	InstanceIds []*string `json:"instanceIds,omitempty"`
}
