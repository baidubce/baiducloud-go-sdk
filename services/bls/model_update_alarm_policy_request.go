package bls

type UpdateAlarmPolicyRequest struct {
	Name                 *string             `json:"name,omitempty"`
	Objects              []*LogStore         `json:"objects,omitempty"`
	Targets              []*Target           `json:"targets,omitempty"`
	TriggerConditions    []*TriggerCondition `json:"triggerConditions,omitempty"`
	Groups               []*string           `json:"groups,omitempty"`
	Schedule             *Schedule           `json:"schedule,omitempty"`
	PendingCount         *int32              `json:"pendingCount,omitempty"`
	RepeatIntervalMinute *int32              `json:"repeatIntervalMinute,omitempty"`
	RecoverWithoutNotice *bool               `json:"recoverWithoutNotice,omitempty"`
	State                *string             `json:"state,omitempty"`
	NoticeState          *string             `json:"noticeState,omitempty"`
	Notices              []*Notice           `json:"notices,omitempty"`
	NoticeRawLogs        []*NoticeRawLog     `json:"noticeRawLogs,omitempty"`
}
