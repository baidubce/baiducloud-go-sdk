package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMongodbSlowLogTrendResponse struct {
	bce.BaseResponse
	Data []*SlowLogTrend `json:"data,omitempty"`
}
