package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListInstancesResponse struct {
	bce.BaseResponse
	OrderBy    *string            `json:"orderBy,omitempty"`
	Order      *string            `json:"order,omitempty"`
	PageNo     *int32             `json:"pageNo,omitempty"`
	PageSize   *int32             `json:"pageSize,omitempty"`
	TotalCount *int32             `json:"totalCount,omitempty"`
	Instances  []*MonitorInstance `json:"instances,omitempty"`
}
