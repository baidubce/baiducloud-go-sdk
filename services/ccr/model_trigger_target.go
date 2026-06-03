package ccr

type TriggerTarget struct {
	Address *string      `json:"address,omitempty"`
	Headers *interface{} `json:"headers,omitempty"`
}
