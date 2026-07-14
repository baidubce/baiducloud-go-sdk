package as

type ScaleCondition struct {
	TargetType         *string `json:"targetType,omitempty"`
	TargetId           *string `json:"targetId,omitempty"`
	Indicator          *string `json:"indicator,omitempty"`
	Threshold          *string `json:"threshold,omitempty"`
	Unit               *string `json:"unit,omitempty"`
	ComparisonOperator *string `json:"comparisonOperator,omitempty"`
	CronTime           *string `json:"cronTime,omitempty"`
	TriggerTime        *string `json:"triggerTime,omitempty"`
	AsType             *string `json:"type,omitempty"`
	PeriodType         *string `json:"periodType,omitempty"`
	PeriodValue        *int32  `json:"periodValue,omitempty"`
	AlarmRule          *string `json:"alarmRule,omitempty"`
}
