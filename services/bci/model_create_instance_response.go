package bci

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateInstanceResponse struct {
	bce.BaseResponse
	InstanceId *string `json:"instanceId,omitempty"`
}
