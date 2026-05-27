package vpc

type ListRateLimitRuleRequest struct {
	GatewayId *string `json:"-"`
	Marker    *string `json:"-"`
	MaxKeys   *int32  `json:"-"`
}
