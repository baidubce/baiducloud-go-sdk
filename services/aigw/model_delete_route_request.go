package aigw

type DeleteRouteRequest struct {
	InstanceId *string `json:"-"`
	RouteName  *string `json:"-"`
	XRegion    *string `json:"-"`
}
