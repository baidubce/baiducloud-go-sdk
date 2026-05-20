package vpc

type CreateDedicatedGatewayHealthCheckRequest struct {
	EtGatewayId           *string `json:"-"`
	ClientToken           *string `json:"-"`
	HealthCheckSourceIp   *string `json:"healthCheckSourceIp,omitempty"`
	HealthCheckType       *string `json:"healthCheckType,omitempty"`
	HealthCheckPort       *int32  `json:"healthCheckPort,omitempty"`
	HealthCheckInterval   *int32  `json:"healthCheckInterval,omitempty"`
	HealthThreshold       *int32  `json:"healthThreshold,omitempty"`
	UnhealthThreshold     *int32  `json:"unhealthThreshold,omitempty"`
	AutoGenerateRouteRule *bool   `json:"autoGenerateRouteRule,omitempty"`
}
