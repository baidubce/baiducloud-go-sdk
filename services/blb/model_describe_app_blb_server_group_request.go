package blb

type DescribeAppBlbServerGroupRequest struct {
	BlbId        *string `json:"-"`
	Name         *string `json:"-"`
	ExactlyMatch *bool   `json:"-"`
	Marker       *string `json:"-"`
	MaxKeys      *int32  `json:"-"`
}
