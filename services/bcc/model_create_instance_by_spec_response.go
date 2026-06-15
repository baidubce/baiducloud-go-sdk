package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateInstanceBySpecResponse struct {
	bce.BaseResponse
	OrderId     *string   `json:"orderId,omitempty"`
	InstanceIds []*string `json:"instanceIds,omitempty"`
	WarningList []*string `json:"warningList,omitempty"`
}
