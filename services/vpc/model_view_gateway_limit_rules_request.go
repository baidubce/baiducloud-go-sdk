package vpc

type ViewGatewayLimitRulesRequest struct {
	ServiceType *string `json:"-"`
	Name        *string `json:"-"`
	GlrId       *string `json:"-"`
	ResourceId  *string `json:"-"`
	Marker      *string `json:"-"`
	MaxKeys     *string `json:"-"`
}
