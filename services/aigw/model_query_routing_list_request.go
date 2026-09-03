package aigw

type QueryRoutingListRequest struct {
	InstanceId *string `json:"-"`
	RouteName  *string `json:"-"`
	PageNo     *int32  `json:"-"`
	PageSize   *int32  `json:"-"`
	OrderBy    *string `json:"-"`
	Order      *string `json:"-"`
	XRegion    *string `json:"-"`
}
