package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListProjectsResponse struct {
	bce.BaseResponse
	PageNo   *int32     `json:"pageNo,omitempty"`
	PageSize *int32     `json:"pageSize,omitempty"`
	Projects []*Project `json:"projects,omitempty"`
	Total    *int32     `json:"total,omitempty"`
}
