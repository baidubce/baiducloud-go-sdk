package aihc

type AlertConfig struct {
	InstanceId   *string   `json:"instanceId,omitempty"`
	AlertItems   []*string `json:"alertItems,omitempty"`
	AihcFor      *string   `json:"for,omitempty"`
	NotifyRuleId *string   `json:"notifyRuleId,omitempty"`
}
