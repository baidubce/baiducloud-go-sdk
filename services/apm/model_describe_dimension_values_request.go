package apm

type DescribeDimensionValuesRequest struct {
	Action        *string   `json:"-"`
	BeginDatetime *string   `json:"beginDatetime,omitempty"`
	EndDatetime   *string   `json:"endDatetime,omitempty"`
	DimensionKey  *string   `json:"dimensionKey,omitempty"`
	Filters       []*Filter `json:"filters,omitempty"`
}
