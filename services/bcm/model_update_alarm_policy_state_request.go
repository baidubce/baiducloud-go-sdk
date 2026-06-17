package bcm

type UpdateAlarmPolicyStateRequest struct {
	Ids   []*string `json:"ids,omitempty"`
	State *string   `json:"state,omitempty"`
}
