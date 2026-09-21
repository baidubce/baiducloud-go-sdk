package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMysqlDatabaseSpaceResponse struct {
	bce.BaseResponse
	Items      []*MysqlDatabaseSpaceModel `json:"items,omitempty"`
	TotalCount *int32                     `json:"totalCount,omitempty"`
}
