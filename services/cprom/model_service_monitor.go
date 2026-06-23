package cprom

type ServiceMonitor struct {
	ApiVersion *string             `json:"apiVersion,omitempty"`
	Kind       *string             `json:"kind,omitempty"`
	Metadata   *ObjectMeta         `json:"metadata,omitempty"`
	Spec       *ServiceMonitorSpec `json:"spec,omitempty"`
}
