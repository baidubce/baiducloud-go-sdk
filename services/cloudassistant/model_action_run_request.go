package cloudassistant

type ActionRunRequest struct {
	Locale             *string            `json:"-"`
	Action             *ActionRef         `json:"action,omitempty"`
	Parameters         *map[string]string `json:"parameters,omitempty"`
	TargetSelectorType *string            `json:"targetSelectorType,omitempty"`
	Targets            []*Target          `json:"targets,omitempty"`
	TargetSelector     *TargetSelector    `json:"targetSelector,omitempty"`
}
