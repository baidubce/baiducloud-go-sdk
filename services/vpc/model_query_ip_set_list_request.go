package vpc

type QueryIpSetListRequest struct {
	IpVersion *string `json:"-"`
	Marker    *string `json:"-"`
	MaxKeys   *int32  `json:"-"`
}
