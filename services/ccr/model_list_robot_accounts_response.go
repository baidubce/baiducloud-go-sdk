package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListRobotAccountsResponse struct {
	bce.BaseResponse
	PageNo   *int32   `json:"pageNo,omitempty"`
	PageSize *int32   `json:"pageSize,omitempty"`
	Total    *int32   `json:"total,omitempty"`
	Robots   []*Robot `json:"robots,omitempty"`
}
