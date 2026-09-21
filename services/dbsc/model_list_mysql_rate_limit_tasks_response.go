package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListMysqlRateLimitTasksResponse struct {
	bce.BaseResponse
	TotalCount *int32             `json:"totalCount,omitempty"`
	Items      []*MySQLFilterItem `json:"items,omitempty"`
}
