package bcm

type UpdateAlarmPolicyNotifyEnabledRequest struct {
	Ids           []*string `json:"ids,omitempty"`
	NotifyEnabled *bool     `json:"notifyEnabled,omitempty"`
}
