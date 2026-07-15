package cloudassistant

type CommandFilter struct {
	Scope              *string `json:"scope,omitempty"`
	Name               *string `json:"name,omitempty"`
	CloudassistantType *string `json:"type,omitempty"`
}
