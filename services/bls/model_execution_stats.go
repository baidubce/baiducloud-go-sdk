package bls

type ExecutionStats struct {
	PolicyId             *string     `json:"policyId,omitempty"`
	PolicyName           *string     `json:"policyName,omitempty"`
	Objects              []*LogStore `json:"objects,omitempty"`
	PendingCount         *int32      `json:"pendingCount,omitempty"`
	RepeatIntervalMinute *int32      `json:"repeatIntervalMinute,omitempty"`
	Notices              []*Notice   `json:"notices,omitempty"`
	TotalCount           *int32      `json:"totalCount,omitempty"`
	FailCount            *int32      `json:"failCount,omitempty"`
	NoticeTotalCount     *int32      `json:"noticeTotalCount,omitempty"`
	NoticeFailCount      *int32      `json:"noticeFailCount,omitempty"`
}
