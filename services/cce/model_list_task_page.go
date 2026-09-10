package cce

type ListTaskPage struct {
	PageNo     *int32         `json:"pageNo,omitempty"`
	PageSize   *int32         `json:"pageSize,omitempty"`
	TotalCount *int32         `json:"totalCount,omitempty"`
	Items      []*interface{} `json:"items,omitempty"`
}
