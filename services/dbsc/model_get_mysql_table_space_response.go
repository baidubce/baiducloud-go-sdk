package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMysqlTableSpaceResponse struct {
	bce.BaseResponse
	Items      []*MysqlTableSpaceModel `json:"items,omitempty"`
	TotalCount *int32                  `json:"totalCount,omitempty"`
}
