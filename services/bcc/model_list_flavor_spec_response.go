package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListFlavorSpecResponse struct {
	bce.BaseResponse
	ZoneResources []*ZoneResourceDetailSpec `json:"zoneResources,omitempty"`
}
