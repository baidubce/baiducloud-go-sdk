package rapidfs

type RemoveCacheNodesRequest struct {
	Action               *string   `json:"-"`
	ClientToken          *string   `json:"-"`
	InstanceId           *string   `json:"instanceId,omitempty"`
	CacheNodeIds         []*string `json:"cacheNodeIds,omitempty"`
	ForceRemoveOnOffline *bool     `json:"forceRemoveOnOffline,omitempty"`
}
