package blb

type DescribeBlbServersRequest struct {
	BlbId   *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
