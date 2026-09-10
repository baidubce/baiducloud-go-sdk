package cce

type CreateNodesClusterExpansionV2Request struct {
	ClusterID   *string        `json:"-"`
	RequestBody []*InstanceSet `json:"无（RequestBody 为数组）,omitempty"`
}
