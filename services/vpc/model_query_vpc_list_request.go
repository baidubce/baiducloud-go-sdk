package vpc

type QueryVpcListRequest struct {
	Marker    *string `json:"-"`
	MaxKeys   *int32  `json:"-"`
	IsDefault *bool   `json:"-"`
	VpcIds    *string `json:"-"`
}
