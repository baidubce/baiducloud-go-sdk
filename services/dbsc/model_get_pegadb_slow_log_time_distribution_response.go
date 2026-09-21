package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetPegadbSlowLogTimeDistributionResponse struct {
	bce.BaseResponse
	Stats []*SCSSlowlogStatsDuartion `json:"stats,omitempty"`
}
