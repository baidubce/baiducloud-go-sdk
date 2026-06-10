package bcc

type AutoSnapshotPolicyModel struct {
	Id              *string  `json:"id,omitempty"`
	Name            *string  `json:"name,omitempty"`
	TimePoints      []*int32 `json:"timePoints,omitempty"`
	RepeatWeekdays  []*int32 `json:"repeatWeekdays,omitempty"`
	Status          *string  `json:"status,omitempty"`
	RetentionDays   *int32   `json:"retentionDays,omitempty"`
	CreatedTime     *string  `json:"createdTime,omitempty"`
	UpdatedTime     *string  `json:"updatedTime,omitempty"`
	DeletedTime     *string  `json:"deletedTime,omitempty"`
	LastExecuteTime *string  `json:"lastExecuteTime,omitempty"`
	VolumeCount     *int32   `json:"volumeCount,omitempty"`
}
