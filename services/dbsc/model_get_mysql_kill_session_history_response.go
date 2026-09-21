package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMysqlKillSessionHistoryResponse struct {
	bce.BaseResponse
	TotalCount *int32                `json:"totalCount,omitempty"`
	Items      []*SessionKillHistory `json:"items,omitempty"`
}
