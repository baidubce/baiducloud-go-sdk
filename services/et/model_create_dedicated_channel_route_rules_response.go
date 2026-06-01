package et

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateDedicatedChannelRouteRulesResponse struct {
	bce.BaseResponse
	RouteRuleId *string `json:"routeRuleId,omitempty"`
}
