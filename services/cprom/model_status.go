package cprom

type Status struct {
	Phase   *string `json:"phase,omitempty"`
	Ready   *bool   `json:"ready,omitempty"`
	Message *string `json:"message,omitempty"`
}
