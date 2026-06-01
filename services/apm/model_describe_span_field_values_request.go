package apm

type DescribeSpanFieldValuesRequest struct {
	Action        *string   `json:"-"`
	BeginDatetime *string   `json:"beginDatetime,omitempty"`
	EndDatetime   *string   `json:"endDatetime,omitempty"`
	Key           *string   `json:"key,omitempty"`
	Filters       []*Filter `json:"filters,omitempty"`
}
