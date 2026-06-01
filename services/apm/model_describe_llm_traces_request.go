package apm

type DescribeLLMTracesRequest struct {
	BeginDatetime *string   `json:"beginDatetime,omitempty"`
	EndDatetime   *string   `json:"endDatetime,omitempty"`
	Filters       []*Filter `json:"filters,omitempty"`
	OrderBy       *string   `json:"orderBy,omitempty"`
	Order         *string   `json:"order,omitempty"`
	PageNo        *int32    `json:"pageNo,omitempty"`
	PageSize      *int32    `json:"pageSize,omitempty"`
}
