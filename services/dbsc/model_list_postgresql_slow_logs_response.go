package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListPostgresqlSlowLogsResponse struct {
	bce.BaseResponse
	TotalCount *int64           `json:"totalCount,omitempty"`
	Logs       []*PGSlowLogInfo `json:"logs,omitempty"`
}
