package blb

type CreateAppBlbIpGroupProtocolRequest struct {
	BlbId                       *string `json:"-"`
	ClientToken                 *string `json:"-"`
	IpGroupId                   *string `json:"ipGroupId,omitempty"`
	BlbType                     *string `json:"type,omitempty"`
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
