package snic

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateSnicResponse struct {
	bce.BaseResponse
	Id        *string `json:"id,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
}
