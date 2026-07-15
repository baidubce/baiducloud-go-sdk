package cloudassistant

type ActionRunRequest struct {
	Locale             *string         `json:"-"`
	Action             *interface{}    `json:"action,omitempty"`
	Parameters         *interface{}    `json:"parameters,omitempty"`
	TargetSelectorType *string         `json:"targetSelectorType,omitempty"`
	Targets            []*interface{}  `json:"targets,omitempty"`
	TargetSelector     *TargetSelector `json:"targetSelector,omitempty"`
}
