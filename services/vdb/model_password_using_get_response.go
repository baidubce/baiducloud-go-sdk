package vdb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type PasswordUsingGETResponse struct {
	bce.BaseResponse
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
}
