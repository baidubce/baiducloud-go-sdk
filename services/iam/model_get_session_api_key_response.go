package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetSessionApiKeyResponse struct {
	bce.BaseResponse
	UserId     *string `json:"userId,omitempty"`
	Token      *string `json:"token,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	ExpireTime *string `json:"expireTime,omitempty"`
}
