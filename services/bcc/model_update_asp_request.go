package bcc

type UpdateAspRequest struct {
	AspId          *string   `json:"aspId,omitempty"`
	Name           *string   `json:"name,omitempty"`
	TimePoints     []*string `json:"timePoints,omitempty"`
	RepeatWeekdays []*string `json:"repeatWeekdays,omitempty"`
	RetentionDays  *string   `json:"retentionDays,omitempty"`
}
