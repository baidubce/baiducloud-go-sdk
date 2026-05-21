package vpc

type ActiveStandbySwitchoverRequest struct {
	RouteRuleId *string `json:"-"`
	ClientToken *string `json:"-"`
}
