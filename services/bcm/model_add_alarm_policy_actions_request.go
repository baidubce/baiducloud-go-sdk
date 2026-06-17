package bcm

type AddAlarmPolicyActionsRequest struct {
	Id      *string         `json:"id,omitempty"`
	Actions []*PolicyAction `json:"actions,omitempty"`
}
