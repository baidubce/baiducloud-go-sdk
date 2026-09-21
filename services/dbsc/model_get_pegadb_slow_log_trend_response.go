package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetPegadbSlowLogTrendResponse struct {
	bce.BaseResponse
	Data []*SCSSlowlogTrend `json:"data,omitempty"`
}
