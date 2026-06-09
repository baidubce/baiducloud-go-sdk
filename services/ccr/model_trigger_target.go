package ccr

type TriggerTarget struct {
	Address *string            `json:"address,omitempty"`
	Headers *map[string]string `json:"headers,omitempty"`
}
