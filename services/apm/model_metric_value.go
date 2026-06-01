package apm

type MetricValue struct {
	Value              *float32 `json:"value,omitempty"`
	CompareToYesterday *float32 `json:"compareToYesterday,omitempty"`
}
