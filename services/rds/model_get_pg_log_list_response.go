package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetPgLogListResponse struct {
	bce.BaseResponse
	PgLogs []*PgLogResponse `json:"pgLogs,omitempty"`
}
