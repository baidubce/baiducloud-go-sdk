package apm

type AlarmEventDetail struct {
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
