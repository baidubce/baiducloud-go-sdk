package vdb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AccountListUsingGETResponse struct {
	bce.BaseResponse
	Usernames []*Account `json:"usernames,omitempty"`
}
