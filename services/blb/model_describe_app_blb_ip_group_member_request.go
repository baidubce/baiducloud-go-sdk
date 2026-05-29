package blb

type DescribeAppBlbIpGroupMemberRequest struct {
	BlbId     *string `json:"-"`
	IpGroupId *string `json:"-"`
	Marker    *string `json:"-"`
	MaxKeys   *int32  `json:"-"`
}
