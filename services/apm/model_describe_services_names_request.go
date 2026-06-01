package apm

type DescribeServicesNamesRequest struct {
	BeginDatetime *string         `json:"beginDatetime,omitempty"`
	EndDatetime   *string         `json:"endDatetime,omitempty"`
	ServiceName   *string         `json:"serviceName,omitempty"`
	ServiceId     *string         `json:"serviceId,omitempty"`
	Language      *string         `json:"language,omitempty"`
	Env           *string         `json:"env,omitempty"`
	Source        *string         `json:"source,omitempty"`
	Tag           *Tag            `json:"tag,omitempty"`
	OrderBy       *string         `json:"orderBy,omitempty"`
	Order         *string         `json:"order,omitempty"`
	MetricFilters []*MetricFilter `json:"metricFilters,omitempty"`
}
