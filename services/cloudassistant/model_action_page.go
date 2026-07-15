package cloudassistant

type ActionPage struct {
	PageNo     *int32    `json:"pageNo,omitempty"`
	PageSize   *int32    `json:"pageSize,omitempty"`
	TotalCount *int32    `json:"totalCount,omitempty"`
	Data       []*Action `json:"data,omitempty"`
}
