package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type TargetRegionMappingListApiResponse struct {
	bce.BaseResponse
	Regions []*string `json:"regions,omitempty"`
}
