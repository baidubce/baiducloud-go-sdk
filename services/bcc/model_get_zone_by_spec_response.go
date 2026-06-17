package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetZoneBySpecResponse struct {
	bce.BaseResponse
	ZoneNames []*string `json:"zoneNames,omitempty"`
}
