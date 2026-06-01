package apm

type DescribeLLMSessionsStatisticsRequest struct {
	BeginDatetime *string   `json:"beginDatetime,omitempty"`
	EndDatetime   *string   `json:"endDatetime,omitempty"`
	Filters       []*Filter `json:"filters,omitempty"`
}
