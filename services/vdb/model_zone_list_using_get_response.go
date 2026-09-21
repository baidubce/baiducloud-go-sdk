package vdb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ZoneListUsingGETResponse struct {
	bce.BaseResponse
	Zones []*ZoneDetail `json:"zones,omitempty"`
}
