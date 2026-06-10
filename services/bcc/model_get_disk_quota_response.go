package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetDiskQuotaResponse struct {
	bce.BaseResponse
	CdsTotalCapacityGB *string     `json:"cdsTotalCapacityGB,omitempty"`
	CdsUsedCapacityGB  *string     `json:"cdsUsedCapacityGB,omitempty"`
	DiskInfos          []*DiskInfo `json:"diskInfos,omitempty"`
}
