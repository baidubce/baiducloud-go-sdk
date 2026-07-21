package as

type AlarmRule struct {
	Id                    *int32       `json:"id,omitempty"`
	Index                 *int32       `json:"index,omitempty"`
	Metric                *string      `json:"metric,omitempty"`
	PeriodInSecond        *int32       `json:"periodInSecond,omitempty"`
	Statistics            *string      `json:"statistics,omitempty"`
	Threshold             *string      `json:"threshold,omitempty"`
	ComparisonOperator    *string      `json:"comparisonOperator,omitempty"`
	EvaluationPeriodCount *int32       `json:"evaluationPeriodCount,omitempty"`
	MetricDimensions      []*Dimension `json:"metricDimensions,omitempty"`
}
