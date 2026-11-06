package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ErrorLogListResponse struct {
	bce.BaseResponse
	Errorlogs []*OpenapiErrorlog `json:"errorlogs,omitempty"`
}
