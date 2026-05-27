package blb

type CreateAppBlbServerGroupPortRequest struct {
	BlbId                       *string `json:"-"`
	ClientToken                 *string `json:"-"`
	SgId                        *string `json:"sgId,omitempty"`
	Port                        *int32  `json:"port,omitempty"`
	BlbType                     *string `json:"type,omitempty"`
	EnableHealthCheck           *bool   `json:"enableHealthCheck,omitempty"`
	HealthCheck                 *string `json:"healthCheck,omitempty"`
	HealthCheckPort             *int32  `json:"healthCheckPort,omitempty"`
	HealthCheckUrlPath          *string `json:"healthCheckUrlPath,omitempty"`
	HealthCheckTimeoutInSecond  *int32  `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckIntervalInSecond *int32  `json:"healthCheckIntervalInSecond,omitempty"`
	HealthCheckDownRetry        *int32  `json:"healthCheckDownRetry,omitempty"`
	HealthCheckUpRetry          *int32  `json:"healthCheckUpRetry,omitempty"`
	HealthCheckNormalStatus     *string `json:"healthCheckNormalStatus,omitempty"`
	HealthCheckHost             *string `json:"healthCheckHost,omitempty"`
	UdpHealthCheckString        *string `json:"udpHealthCheckString,omitempty"`
}
