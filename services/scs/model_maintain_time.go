package scs

type MaintainTime struct {
	StartTime *string   `json:"startTime,omitempty"`
	Duration  *int32    `json:"duration,omitempty"`
	Period    []*string `json:"period,omitempty"`
}
