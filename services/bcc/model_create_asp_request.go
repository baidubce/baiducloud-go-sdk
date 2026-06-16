package bcc

type CreateAspRequest struct {
	Name           *string  `json:"name,omitempty"`
	TimePoints     []*int32 `json:"timePoints,omitempty"`
	RepeatWeekdays []*int32 `json:"repeatWeekdays,omitempty"`
	RetentionDays  *string  `json:"retentionDays,omitempty"`
}
