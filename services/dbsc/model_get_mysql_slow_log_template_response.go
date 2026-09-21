package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMysqlSlowLogTemplateResponse struct {
	bce.BaseResponse
	Items      []*MySQLSlowLogTemplate `json:"items,omitempty"`
	TotalCount *int32                  `json:"totalCount,omitempty"`
}
