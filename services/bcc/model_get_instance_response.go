package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetInstanceResponse struct {
	bce.BaseResponse
	Instance *InstanceModel `json:"instance,omitempty"`
}
