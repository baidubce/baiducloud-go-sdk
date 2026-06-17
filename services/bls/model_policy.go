package bls

type Policy struct {
	Name                 *string             `json:"name,omitempty"`
	Id                   *string             `json:"id,omitempty"`
	State                *string             `json:"state,omitempty"`
	NoticeState          *string             `json:"noticeState,omitempty"`
	CreatedTime          *string             `json:"createdTime,omitempty"`
	UpdatedTime          *string             `json:"updatedTime,omitempty"`
	Objects              []*LogStore         `json:"objects,omitempty"`
	Targets              []*Target           `json:"targets,omitempty"`
	TriggerConditions    []*TriggerCondition `json:"triggerConditions,omitempty"`
	Groups               []*string           `json:"groups,omitempty"`
	Schedule             *Schedule           `json:"schedule,omitempty"`
	PendingCount         *int32              `json:"pendingCount,omitempty"`
	RepeatIntervalMinute *int32              `json:"repeatIntervalMinute,omitempty"`
	RecoverAlarmNotice   *bool               `json:"recoverAlarmNotice,omitempty"`
	Notices              []*Notice           `json:"notices,omitempty"`
	NoticeWithRawLog     *bool               `json:"noticeWithRawLog,omitempty"`
	NoticeRawConfigs     []*NoticeRawLog     `json:"noticeRawConfigs,omitempty"`
}
