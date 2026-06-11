package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListOsResponse struct {
	bce.BaseResponse
	OsInfo []*OsModel `json:"osInfo,omitempty"`
}
