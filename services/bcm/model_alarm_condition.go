package bcm

type AlarmCondition struct {
	MetricName       *string            `json:"metricName,omitempty"`
	MetricDimensions []*MetricDimension `json:"metricDimensions,omitempty"`
	Operator         *string            `json:"operator,omitempty"`
	Threshold        *float32           `json:"threshold,omitempty"`
	AggregateType    *string            `json:"aggregateType,omitempty"`
	WindowSeconds    *int32             `json:"windowSeconds,omitempty"`
	DisplayUnit      *string            `json:"displayUnit,omitempty"`
	DisplayThreshold *string            `json:"displayThreshold,omitempty"`
}
