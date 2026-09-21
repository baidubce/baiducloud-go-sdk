package vdb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type Getinstancelistusingget1Response struct {
	bce.BaseResponse
	Instances []*Instance `json:"instances,omitempty"`
}
