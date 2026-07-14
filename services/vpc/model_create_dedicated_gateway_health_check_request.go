package vpc

type CreateDedicatedGatewayHealthCheckRequest struct {
	EtGatewayId           *string `json:"-"`
	ClientToken           *string `json:"-"`
	DcphyId               *string `json:"dcphyId,omitempty"`
	ChannelId             *string `json:"channelId,omitempty"`
	SubnetId              *string `json:"subnetId,omitempty"`
	HealthCheckSourceIp   *string `json:"healthCheckSourceIp,omitempty"`
	HealthCheckType       *string `json:"healthCheckType,omitempty"`
	HealthCheckInterval   *int32  `json:"healthCheckInterval,omitempty"`
	HealthThreshold       *int32  `json:"healthThreshold,omitempty"`
	UnhealthThreshold     *int32  `json:"unhealthThreshold,omitempty"`
	AutoGenerateRouteRule *bool   `json:"autoGenerateRouteRule,omitempty"`
}
