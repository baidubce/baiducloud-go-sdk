package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetInstanceUserDataInfoResponse struct {
	bce.BaseResponse
	UserData   *string `json:"userData,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
}
