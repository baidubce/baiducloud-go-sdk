package rapidfs

type DescribeInstancesRequest struct {
	Action  *string   `json:"-"`
	Filters []*Filter `json:"filters,omitempty"`
	MaxKeys *int32    `json:"maxKeys,omitempty"`
	Marker  *string   `json:"marker,omitempty"`
}
