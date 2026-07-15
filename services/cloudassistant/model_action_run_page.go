package cloudassistant

type ActionRunPage struct {
	PageNo     *int32       `json:"pageNo,omitempty"`
	PageSize   *int32       `json:"pageSize,omitempty"`
	TotalCount *int32       `json:"totalCount,omitempty"`
	Data       []*ActionRun `json:"data,omitempty"`
}
