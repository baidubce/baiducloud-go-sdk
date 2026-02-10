package vpc

type QueryTheReservedNetworkSegmentListRequest struct {
	SubnetId *string `json:"-"`
	Marker   *string `json:"-"`
	MaxKeys  *int32  `json:"-"`
}
