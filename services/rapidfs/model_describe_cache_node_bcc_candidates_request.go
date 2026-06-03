package rapidfs

type DescribeCacheNodeBccCandidatesRequest struct {
	Action     *string   `json:"-"`
	InstanceId *string   `json:"instanceId,omitempty"`
	VpcId      *string   `json:"vpcId,omitempty"`
	Filters    []*Filter `json:"filters,omitempty"`
	MaxKeys    *int32    `json:"maxKeys,omitempty"`
	Marker     *string   `json:"marker,omitempty"`
}
