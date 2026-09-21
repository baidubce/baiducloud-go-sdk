package vdb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateInstanceUsingPOSTResponse struct {
	bce.BaseResponse
	EnableEncryption *bool     `json:"enableEncryption,omitempty"`
	InstanceIdList   []*string `json:"instanceIdList,omitempty"`
	OrderId          *string   `json:"orderId,omitempty"`
}
