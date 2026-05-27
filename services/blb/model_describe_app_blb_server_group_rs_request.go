package blb

type DescribeAppBlbServerGroupRsRequest struct {
	BlbId   *string `json:"-"`
	SgId    *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
