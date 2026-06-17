package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeAlarmMaskingResponse struct {
	bce.BaseResponse
	Id                  *string           `json:"id,omitempty"`
	CreatedTime         *string           `json:"createdTime,omitempty"`
	UpdatedTime         *string           `json:"updatedTime,omitempty"`
	State               *string           `json:"state,omitempty"`
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
