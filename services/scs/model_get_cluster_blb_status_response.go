package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetClusterBlbStatusResponse struct {
	bce.BaseResponse
	Alive  *bool   `json:"alive,omitempty"`
	Errors *string `json:"errors,omitempty"`
}
