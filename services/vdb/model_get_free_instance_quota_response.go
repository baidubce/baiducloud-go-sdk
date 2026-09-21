package vdb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetFreeInstanceQuotaResponse struct {
	bce.BaseResponse
	FreeQuota *int32 `json:"freeQuota,omitempty"`
}
