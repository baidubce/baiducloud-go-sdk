package rds

type UpdateTimeWindowRequest struct {
	InstanceId        *string `json:"-"`
	MaintainStartTime *string `json:"maintainStartTime,omitempty"`
	MaintainDuration  *int32  `json:"maintainDuration,omitempty"`
}
