package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetPostgresqlSlowLogTimeDistributionResponse struct {
	bce.BaseResponse
	Stats []*StatsByRangeItem `json:"stats,omitempty"`
}
