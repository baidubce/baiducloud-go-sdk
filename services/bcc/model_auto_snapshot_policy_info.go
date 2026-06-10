package bcc

type AutoSnapshotPolicyInfo struct {
	Id             *string  `json:"id,omitempty"`
	Name           *string  `json:"name,omitempty"`
	TimePoints     []*int32 `json:"timePoints,omitempty"`
	RepeatWeekdays []*int32 `json:"repeatWeekdays,omitempty"`
	RetentionDays  *int32   `json:"retentionDays,omitempty"`
	Status         *string  `json:"status,omitempty"`
}
