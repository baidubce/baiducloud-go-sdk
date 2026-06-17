package bcm

type DeleteAlarmPolicyActionsRequest struct {
	Id      *string         `json:"id,omitempty"`
	Actions []*PolicyAction `json:"actions,omitempty"`
}
