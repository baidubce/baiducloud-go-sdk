package vpc

type QueryRoutingTableRequest struct {
	RouteTableId *string `json:"-"`
	VpcId        *string `json:"-"`
}
