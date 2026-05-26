package vpc

type ListPeerConnRequest struct {
	VpcId   *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
