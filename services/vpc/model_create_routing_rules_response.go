package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateRoutingRulesResponse struct {
	bce.BaseResponse
	RouteRuleId  *string   `json:"routeRuleId,omitempty"`
	RouteRuleIds []*string `json:"routeRuleIds,omitempty"`
}
