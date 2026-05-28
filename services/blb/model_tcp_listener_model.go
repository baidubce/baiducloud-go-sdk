package blb

type TCPListenerModel struct {
	ListenerPort               *int32  `json:"listenerPort,omitempty"`
	BackendPort                *int32  `json:"backendPort,omitempty"`
	Scheduler                  *string `json:"scheduler,omitempty"`
	HealthCheckTimeoutInSecond *int32  `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        *int32  `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         *int32  `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           *int32  `json:"healthyThreshold,omitempty"`
	TcpSessionTimeout          *int32  `json:"tcpSessionTimeout,omitempty"`
}
