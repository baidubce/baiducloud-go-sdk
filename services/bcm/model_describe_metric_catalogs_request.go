package bcm

type DescribeMetricCatalogsRequest struct {
	Locale              *string         `json:"-"`
	Scope               *string         `json:"scope,omitempty"`
	ResourceType        *string         `json:"resourceType,omitempty"`
	Catalog             *string         `json:"catalog,omitempty"`
	Filters             []*MetricFilter `json:"filters,omitempty"`
	FiltersKey          *string         `json:"filters[].key,omitempty"`
	FiltersOp           *string         `json:"filters[].op,omitempty"`
	FiltersValue        *string         `json:"filters[].value,omitempty"`
	FiltersValues       []*string       `json:"filters[].values,omitempty"`
	IncludingDimensions []*string       `json:"includingDimensions,omitempty"`
	ExcludingDimensions []*string       `json:"excludingDimensions,omitempty"`
}
