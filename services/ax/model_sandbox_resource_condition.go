package ax

type SandboxResourceCondition struct {
	AxType  *string `json:"type,omitempty"`
	Status  *string `json:"status,omitempty"`
	Reason  *string `json:"reason,omitempty"`
	Message *string `json:"message,omitempty"`
}
