package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetSubnetListResponse struct {
	bce.BaseResponse
	Subnets []*Subnet `json:"subnets,omitempty"`
}
