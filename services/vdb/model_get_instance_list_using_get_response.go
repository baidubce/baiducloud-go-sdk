package vdb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetInstanceListUsingGETResponse struct {
	bce.BaseResponse
	Instances []*Instance `json:"instances,omitempty"`
}
