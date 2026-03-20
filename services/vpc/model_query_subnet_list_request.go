package vpc

type QuerySubnetListRequest struct {
	Marker     *string `json:"-"`
	MaxKeys    *int32  `json:"-"`
	VpcId      *string `json:"-"`
	ZoneName   *string `json:"-"`
	SubnetType *string `json:"-"`
	SubnetIds  *string `json:"-"`
}
