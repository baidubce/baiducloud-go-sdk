package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryTheListOfAvailableZonesResponse struct {
	bce.BaseResponse
	Zones []*ZoneDetail `json:"zones,omitempty"`
}
