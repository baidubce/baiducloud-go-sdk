package blb

type CreateBlbTcpListenerRequest struct {
	BlbId                      *string `json:"-"`
	ClientToken                *string `json:"-"`
	ListenerPort               *int32  `json:"listenerPort,omitempty"`
	BackendPort                *int32  `json:"backendPort,omitempty"`
	Scheduler                  *string `json:"scheduler,omitempty"`
	TcpSessionTimeout          *int32  `json:"tcpSessionTimeout,omitempty"`
	HealthCheckType            *string `json:"healthCheckType,omitempty"`
	HealthCheckTimeoutInSecond *int32  `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        *int32  `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         *int32  `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           *int32  `json:"healthyThreshold,omitempty"`
}
