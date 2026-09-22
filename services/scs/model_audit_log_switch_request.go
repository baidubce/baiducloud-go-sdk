package scs

type AuditLogSwitchRequest struct {
	InstanceId *string `json:"-"`
	Action     *string `json:"action,omitempty"`
}
