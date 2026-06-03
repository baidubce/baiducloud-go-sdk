package rapidfs

type DescribeSpecsRequest struct {
	Filters []*Filter `json:"filters,omitempty"`
}
