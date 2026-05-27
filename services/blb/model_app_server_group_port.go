package blb

type AppServerGroupPort struct {
	Id                          *string `json:"id,omitempty"`
	Port                        *int32  `json:"port,omitempty"`
	BlbType                     *string `json:"type,omitempty"`
	Status                      *string `json:"status,omitempty"`
	HealthCheck                 *string `json:"healthCheck,omitempty"`
	HealthCheckPort             *int32  `json:"healthCheckPort,omitempty"`
	HealthCheckTimeoutInSecond  *int32  `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckIntervalInSecond *int32  `json:"healthCheckIntervalInSecond,omitempty"`
	HealthCheckDownRetry        *int32  `json:"healthCheckDownRetry,omitempty"`
	HealthCheckUpRetry          *int32  `json:"healthCheckUpRetry,omitempty"`
	HealthCheckNormalStatus     *string `json:"healthCheckNormalStatus,omitempty"`
	HealthCheckUrlPath          *string `json:"healthCheckUrlPath,omitempty"`
	HealthCheckHost             *string `json:"healthCheckHost,omitempty"`
	UdpHealthCheckString        *string `json:"udpHealthCheckString,omitempty"`
}
