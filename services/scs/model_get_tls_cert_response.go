package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetTlsCertResponse struct {
	bce.BaseResponse
	Status     *string `json:"status,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	ExpireTime *string `json:"expireTime,omitempty"`
	Ca         *string `json:"ca,omitempty"`
}
