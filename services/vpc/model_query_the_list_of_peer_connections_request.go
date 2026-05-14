package vpc

type QueryTheListOfPeerConnectionsRequest struct {
	VpcId   *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
