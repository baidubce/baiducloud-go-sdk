package vdb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetQuotaUsingGETResponse struct {
	bce.BaseResponse
	Quota *int64 `json:"quota,omitempty"`
}
