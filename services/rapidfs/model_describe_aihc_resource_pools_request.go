package rapidfs

type DescribeAihcResourcePoolsRequest struct {
	Action  *string `json:"-"`
	VpcId   *string `json:"vpcId,omitempty"`
	MaxKeys *int32  `json:"maxKeys,omitempty"`
	Marker  *string `json:"marker,omitempty"`
}
