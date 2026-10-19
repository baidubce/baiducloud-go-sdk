package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QuerySubnetListResponse struct {
	bce.BaseResponse
	Subnets []*SubnetVo `json:"subnets,omitempty"`
}
