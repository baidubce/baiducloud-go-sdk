package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetPostgresqlSlowLogTrendResponse struct {
	bce.BaseResponse
	Data []*SlowTrendDataPoint `json:"data,omitempty"`
}
