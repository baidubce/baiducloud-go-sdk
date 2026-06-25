package cprom

type PodMonitor struct {
	ApiVersion *string         `json:"apiVersion,omitempty"`
	Kind       *string         `json:"kind,omitempty"`
	Metadata   *ObjectMeta     `json:"metadata,omitempty"`
	Spec       *PodMonitorSpec `json:"spec,omitempty"`
}
