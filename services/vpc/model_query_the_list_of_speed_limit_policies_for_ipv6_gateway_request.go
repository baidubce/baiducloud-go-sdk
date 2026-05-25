package vpc

type QueryTheListOfSpeedLimitPoliciesForIpv6GatewayRequest struct {
	GatewayId *string `json:"-"`
	Marker    *string `json:"-"`
	MaxKeys   *int32  `json:"-"`
}
