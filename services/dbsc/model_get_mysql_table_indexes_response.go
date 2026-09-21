package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMysqlTableIndexesResponse struct {
	bce.BaseResponse
	Indexes []*TableIndexDetailItem `json:"indexes,omitempty"`
}
