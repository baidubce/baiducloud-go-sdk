package apm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ApmDescribeAlarmResponse struct {
	bce.BaseResponse
	Success       *bool              `json:"success,omitempty"`
	Code          *string            `json:"code,omitempty"`
	Message       *string            `json:"message,omitempty"`
	Id            *string            `json:"id,omitempty"`
	StartTime     *string            `json:"startTime,omitempty"`
	EndTime       *string            `json:"endTime,omitempty"`
	Duration      *int32             `json:"duration,omitempty"`
	InitState     *string            `json:"initState,omitempty"`
	State         *string            `json:"state,omitempty"`
	CloseReason   *string            `json:"closeReason,omitempty"`
	CurrentValue  *string            `json:"currentValue,omitempty"`
	MonitorObject *string            `json:"monitorObject,omitempty"`
	Policy        *AlarmPolicyDetail `json:"policy,omitempty"`
}
