package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryRoutingTableResponse struct {
	bce.BaseResponse
	Name         *string      `json:"name,omitempty"`
	RouteTableId *string      `json:"routeTableId,omitempty"`
	VpcId        *string      `json:"vpcId,omitempty"`
	RouteRules   []*RouteRule `json:"routeRules,omitempty"`
}
