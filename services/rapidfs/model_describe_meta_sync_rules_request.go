package rapidfs

type DescribeMetaSyncRulesRequest struct {
	InstanceId *string   `json:"instanceId,omitempty"`
	Filters    []*Filter `json:"filters,omitempty"`
	MaxKeys    *int32    `json:"maxKeys,omitempty"`
	Marker     *string   `json:"marker,omitempty"`
}
