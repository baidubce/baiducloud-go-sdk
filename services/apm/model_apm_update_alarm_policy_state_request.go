package apm

type ApmUpdateAlarmPolicyStateRequest struct {
	Ids   []*string `json:"ids,omitempty"`
	State *string   `json:"state,omitempty"`
}
