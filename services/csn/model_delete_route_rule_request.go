package csn

type DeleteRouteRuleRequest struct {
	CsnRtId     *string `json:"-"`
	CsnRtRuleId *string `json:"-"`
	ClientToken *string `json:"-"`
}
