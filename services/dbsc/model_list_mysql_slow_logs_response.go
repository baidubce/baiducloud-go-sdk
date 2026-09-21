package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListMysqlSlowLogsResponse struct {
	bce.BaseResponse
	Logs       []*MySQLSlowLogDetail `json:"logs,omitempty"`
	TotalCount *int32                `json:"totalCount,omitempty"`
}
