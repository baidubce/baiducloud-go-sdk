package vpc

type QueryIpAddressFamilyListRequest struct {
	IpVersion *string `json:"-"`
	Marker    *string `json:"-"`
	MaxKeys   *int32  `json:"-"`
}
