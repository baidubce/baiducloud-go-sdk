package snic

type ListSnicRequest struct {
	VpcId     *string `json:"-"`
	Name      *string `json:"-"`
	IpAddress *string `json:"-"`
	Status    *string `json:"-"`
	SubnetId  *string `json:"-"`
	Service   *string `json:"-"`
	Marker    *string `json:"-"`
	MaxKeys   *int32  `json:"-"`
}
