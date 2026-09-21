package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetPegadbSlowLogTemplateResponse struct {
	bce.BaseResponse
	TotalCount *int64                   `json:"totalCount,omitempty"`
	Result     []*SCSSlowLogSummaryItem `json:"result,omitempty"`
}
