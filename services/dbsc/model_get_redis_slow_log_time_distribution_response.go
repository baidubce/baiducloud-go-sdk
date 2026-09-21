package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetRedisSlowLogTimeDistributionResponse struct {
	bce.BaseResponse
	Stats []*SCSSlowlogStatsDuartion `json:"stats,omitempty"`
}
