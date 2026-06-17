package bls

type Execution struct {
	Time        *string                 `json:"time,omitempty"`
	State       *string                 `json:"state,omitempty"`
	NoticeState *string                 `json:"noticeState,omitempty"`
	Reason      *string                 `json:"reason,omitempty"`
	Values      *map[string]interface{} `json:"values,omitempty"`
	Notices     []*Notice               `json:"notices,omitempty"`
	RawLogs     []*RawLog               `json:"rawLogs,omitempty"`
}
