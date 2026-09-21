package vdb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListRecordsUsingGETResponse struct {
	bce.BaseResponse
	Records []*Record `json:"records,omitempty"`
	Total   *int64    `json:"total,omitempty"`
}
