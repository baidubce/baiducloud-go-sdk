package vpc

type QueryRoutingRulesRequest struct {
	RouteTableId *string `json:"-"`
	VpcId        *string `json:"-"`
	Marker       *string `json:"-"`
	MaxKeys      *int32  `json:"-"`
}
