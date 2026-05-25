package vpc

type QueryTheListOfHighlyAvailableVirtualIpsRequest struct {
	VpcId   *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
