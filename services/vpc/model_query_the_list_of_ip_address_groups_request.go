package vpc

type QueryTheListOfIpAddressGroupsRequest struct {
	IpVersion *string `json:"-"`
	Marker    *string `json:"-"`
	MaxKeys   *int32  `json:"-"`
}
