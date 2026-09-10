package cce

type ListInstanceGroupPage struct {
	PageNo     *int32         `json:"pageNo,omitempty"`
	PageSize   *int32         `json:"pageSize,omitempty"`
	TotalCount *int32         `json:"totalCount,omitempty"`
	List       []*interface{} `json:"list,omitempty"`
}
