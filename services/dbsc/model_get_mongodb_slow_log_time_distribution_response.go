package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMongodbSlowLogTimeDistributionResponse struct {
	bce.BaseResponse
	Stats []*MySQLSlowLogStatsByDuration `json:"stats,omitempty"`
}
