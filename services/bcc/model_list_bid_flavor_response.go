package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListBidFlavorResponse struct {
	bce.BaseResponse
	ZoneResources []*ZoneResource `json:"zoneResources,omitempty"`
}
