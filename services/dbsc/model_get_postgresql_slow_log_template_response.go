package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetPostgresqlSlowLogTemplateResponse struct {
	bce.BaseResponse
	TotalCount *int64                      `json:"totalCount,omitempty"`
	Items      []*APIPGSlowLogTemplateItem `json:"items,omitempty"`
}
