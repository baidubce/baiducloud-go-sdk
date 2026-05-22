package blb

type ListenerModel struct {
	Port        *string `json:"port,omitempty"`
	BlbType     *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
}
