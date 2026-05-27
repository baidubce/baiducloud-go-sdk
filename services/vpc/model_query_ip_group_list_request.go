package vpc

type QueryIpGroupListRequest struct {
	IpVersion *string `json:"-"`
	Marker    *string `json:"-"`
	MaxKeys   *int32  `json:"-"`
}
