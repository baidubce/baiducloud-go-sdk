package cloudassistant

type BatchGetAgentRequest struct {
	Hosts []*Host `json:"hosts,omitempty"`
}
