package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMysqlSlowLogTimeDistributionResponse struct {
	bce.BaseResponse
	Stats []*MySQLSlowLogStatsByDuration `json:"stats,omitempty"`
}
