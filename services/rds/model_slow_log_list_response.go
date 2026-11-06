package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type SlowLogListResponse struct {
	bce.BaseResponse
	Slowlogs []*OpenapiSlowlog `json:"slowlogs,omitempty"`
}
