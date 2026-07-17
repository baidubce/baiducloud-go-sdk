package cloudassistant

type UpdateActionRequest struct {
	Execution          *string            `json:"execution,omitempty"`
	Action             *Action            `json:"action,omitempty"`
	Parameters         *map[string]string `json:"parameters,omitempty"`
	TargetSelectorType *string            `json:"targetSelectorType,omitempty"`
	Targets            []*Target          `json:"targets,omitempty"`
	TargetSelector     *TargetSelector    `json:"targetSelector,omitempty"`
}
