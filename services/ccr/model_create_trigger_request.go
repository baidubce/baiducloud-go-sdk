package ccr

type CreateTriggerRequest struct {
	InstanceId  *string          `json:"-"`
	Description *string          `json:"description,omitempty"`
	EventTypes  []*string        `json:"eventTypes,omitempty"`
	Filters     []*TriggerFilter `json:"filters,omitempty"`
	Name        *string          `json:"name,omitempty"`
	Targets     []*TriggerTarget `json:"targets,omitempty"`
}
