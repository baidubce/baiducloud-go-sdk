package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMysqlSlowLogTrendResponse struct {
	bce.BaseResponse
	Data []*MySQLSlowLogTrend `json:"data,omitempty"`
}
