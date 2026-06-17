package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeRegionsResponse struct {
	bce.BaseResponse
	Regions []*RegionModel `json:"regions,omitempty"`
}
