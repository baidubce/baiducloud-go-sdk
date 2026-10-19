package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DialTestInterfaceResponse struct {
	bce.BaseResponse
	Alive  *bool   `json:"alive,omitempty"`
	Errors *string `json:"errors,omitempty"`
}
