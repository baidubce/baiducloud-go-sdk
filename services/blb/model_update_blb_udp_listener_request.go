package blb

type UpdateBlbUdpListenerRequest struct {
	BlbId                      *string `json:"-"`
	ClientToken                *string `json:"-"`
	ListenerPort               *int32  `json:"-"`
	BackendPort                *int32  `json:"backendPort,omitempty"`
	Scheduler                  *string `json:"scheduler,omitempty"`
	HealthCheckType            *string `json:"healthCheckType,omitempty"`
	HealthCheckPort            *int32  `json:"healthCheckPort,omitempty"`
	HealthCheckTimeoutInSecond *int32  `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        *int32  `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         *int32  `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           *int32  `json:"healthyThreshold,omitempty"`
	HealthCheckString          *string `json:"healthCheckString,omitempty"`
	UdpSessionTimeout          *int32  `json:"udpSessionTimeout,omitempty"`
}
