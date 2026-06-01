package apm

type AlarmRule struct {
	Operator        *string      `json:"operator,omitempty"`
	Rules           []*AlarmRule `json:"rules,omitempty"`
	Metric          *string      `json:"metric,omitempty"`
	WindowInSeconds *int32       `json:"windowInSeconds,omitempty"`
	Aggregate       *string      `json:"aggregate,omitempty"`
	DisplayValue    *float32     `json:"displayValue,omitempty"`
	DisplayUnit     *string      `json:"displayUnit,omitempty"`
}
