package rapidfs

type DescribeSpecsRequest struct {
	Action  *string   `json:"-"`
	Filters []*Filter `json:"filters,omitempty"`
}
