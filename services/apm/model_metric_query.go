package apm

type MetricQuery struct {
	Name      *string   `json:"name,omitempty"`
	CompareTo []*string `json:"compareTo,omitempty"`
	Filters   []*Filter `json:"filters,omitempty"`
}
