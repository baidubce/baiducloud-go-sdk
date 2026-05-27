package blb

type AppRsPortModel struct {
	ListenerPort        *int32  `json:"listenerPort,omitempty"`
	BackendPort         *string `json:"backendPort,omitempty"`
	PortType            *string `json:"portType,omitempty"`
	HealthCheckPortType *string `json:"healthCheckPortType,omitempty"`
	Status              *string `json:"status,omitempty"`
	PortId              *string `json:"portId,omitempty"`
	PolicyId            *string `json:"policyId,omitempty"`
}
