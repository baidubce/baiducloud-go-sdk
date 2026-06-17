package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListZonesResponse struct {
	bce.BaseResponse
	Zones []*ZoneModel `json:"zones,omitempty"`
}
