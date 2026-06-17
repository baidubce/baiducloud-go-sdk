package bcm

type CreateAlarmMaskingRequest struct {
	Name                *string           `json:"name,omitempty"`
	Scope               *string           `json:"scope,omitempty"`
	ResourceType        *string           `json:"resourceType,omitempty"`
	PolicyId            *string           `json:"policyId,omitempty"`
	Instances           []*TargetInstance `json:"instances,omitempty"`
	Region              *string           `json:"region,omitempty"`
	MetricNames         []*string         `json:"metricNames,omitempty"`
	PeriodType          *string           `json:"periodType,omitempty"`
	BeginTime           *string           `json:"beginTime,omitempty"`
	EndTime             *string           `json:"endTime,omitempty"`
	Tz                  *string           `json:"tz,omitempty"`
	DailyBeginTimestamp *int32            `json:"dailyBeginTimestamp,omitempty"`
	DailyEndTimestamp   *int32            `json:"dailyEndTimestamp,omitempty"`
}
