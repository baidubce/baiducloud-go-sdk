package vpc

type ListEgressOnlyRuleRequest struct {
	GatewayId *string `json:"-"`
	Marker    *string `json:"-"`
	MaxKeys   *int32  `json:"-"`
}
