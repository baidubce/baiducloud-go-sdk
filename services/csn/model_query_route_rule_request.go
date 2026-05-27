package csn

type QueryRouteRuleRequest struct {
	CsnRtId *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
