package blb

type HTTPListenerModel struct {
	ListenerPort               *int32                     `json:"listenerPort,omitempty"`
	BackendPort                *int32                     `json:"backendPort,omitempty"`
	Scheduler                  *string                    `json:"scheduler,omitempty"`
	KeepSession                *bool                      `json:"keepSession,omitempty"`
	KeepSessionType            *string                    `json:"keepSessionType,omitempty"`
	KeepSessionDuration        *int32                     `json:"keepSessionDuration,omitempty"`
	KeepSessionCookieName      *string                    `json:"keepSessionCookieName,omitempty"`
	XForwardFor                *bool                      `json:"xForwardFor,omitempty"`
	AdditionalAttributes       *AdditionalAttributesModel `json:"additionalAttributes,omitempty"`
	HealthCheckType            *string                    `json:"healthCheckType,omitempty"`
	HealthCheckPort            *int32                     `json:"healthCheckPort,omitempty"`
	HealthCheckURI             *string                    `json:"healthCheckURI,omitempty"`
	HealthCheckTimeoutInSecond *int32                     `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        *int32                     `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         *int32                     `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           *int32                     `json:"healthyThreshold,omitempty"`
	HealthCheckNormalStatus    *string                    `json:"healthCheckNormalStatus,omitempty"`
	HealthCheckHost            *string                    `json:"healthCheckHost,omitempty"`
	ServerTimeout              *int32                     `json:"serverTimeout,omitempty"`
	RedirectPort               *int32                     `json:"redirectPort,omitempty"`
}
