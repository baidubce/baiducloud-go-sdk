package apm

type ApmUpdateAlarmPolicyActionRequest struct {
	Id      *string        `json:"id,omitempty"`
	Actions []*AlarmAction `json:"actions,omitempty"`
}
