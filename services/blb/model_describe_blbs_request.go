package blb

type DescribeBlbsRequest struct {
	Address      *string `json:"-"`
	Name         *string `json:"-"`
	BlbId        *string `json:"-"`
	BccId        *string `json:"-"`
	ExactlyMatch *bool   `json:"-"`
	Marker       *string `json:"-"`
	MaxKeys      *int32  `json:"-"`
	Type         *string `json:"-"`
}
