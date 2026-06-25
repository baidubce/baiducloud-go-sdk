package cprom

type CreateServiceMonitorRequest struct {
	InstanceId *string             `json:"-"`
	AgentId    *string             `json:"-"`
	ApiVersion *string             `json:"apiVersion,omitempty"`
	Kind       *string             `json:"kind,omitempty"`
	Metadata   *ObjectMeta         `json:"metadata,omitempty"`
	Spec       *ServiceMonitorSpec `json:"spec,omitempty"`
}
