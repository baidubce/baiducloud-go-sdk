package cprom

type ScopeDetail struct {
	ScopeName    *string `json:"scopeName,omitempty"`
	Humanization *string `json:"humanization,omitempty"`
	DocLink      *string `json:"docLink,omitempty"`
}
