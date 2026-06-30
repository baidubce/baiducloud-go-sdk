package bci

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetInstanceResponse struct {
	bce.BaseResponse
	Instance *InstanceDetailModel `json:"instance,omitempty"`
}
