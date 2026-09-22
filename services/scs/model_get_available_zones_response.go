package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetAvailableZonesResponse struct {
	bce.BaseResponse
	Zones []*ZoneNames `json:"zones,omitempty"`
}
