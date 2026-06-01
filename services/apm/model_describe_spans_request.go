package apm

type DescribeSpansRequest struct {
	Action        *string   `json:"-"`
	BeginDatetime *string   `json:"beginDatetime,omitempty"`
	EndDatetime   *string   `json:"endDatetime,omitempty"`
	Filters       []*Filter `json:"filters,omitempty"`
	OrderBy       *string   `json:"orderBy,omitempty"`
	Order         *string   `json:"order,omitempty"`
	Marker        *string   `json:"marker,omitempty"`
}
