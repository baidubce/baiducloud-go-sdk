package aigw

type RouteSummary struct {
	RouteName             *string          `json:"routeName,omitempty"`
	RouteStatus           *string          `json:"routeStatus,omitempty"`
	Domains               []*string        `json:"domains,omitempty"`
	AssociatedDomainCount *int32           `json:"associatedDomainCount,omitempty"`
	MatchPath             *RouteMatchPath  `json:"matchPath,omitempty"`
	TargetService         []*TargetService `json:"targetService,omitempty"`
	CreateTime            *string          `json:"createTime,omitempty"`
	SrcProduct            *string          `json:"srcProduct,omitempty"`
}
