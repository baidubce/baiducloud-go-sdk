package blb

type DescribeAppBlbPolicyRequest struct {
	BlbId   *string `json:"-"`
	Port    *int32  `json:"-"`
	Type    *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
