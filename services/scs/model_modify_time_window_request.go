package scs

type ModifyTimeWindowRequest struct {
	InstanceId *string  `json:"-"`
	StartTime  *string  `json:"startTime,omitempty"`
	Duration   *int32   `json:"duration,omitempty"`
	Period     []*int32 `json:"period,omitempty"`
}
