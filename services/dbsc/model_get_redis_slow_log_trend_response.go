package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetRedisSlowLogTrendResponse struct {
	bce.BaseResponse
	Data []*SCSSlowlogTrend `json:"data,omitempty"`
}
