package ccr

type ReplicationTrigger struct {
	TriggerSettings *ReplicationSettings `json:"triggerSettings,omitempty"`
	CcrType         *string              `json:"type,omitempty"`
}
