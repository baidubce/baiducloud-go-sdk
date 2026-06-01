package blb

type DescribeServicesRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
