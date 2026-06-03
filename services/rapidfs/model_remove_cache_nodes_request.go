package rapidfs

type RemoveCacheNodesRequest struct {
	ClientToken          *string   `json:"-"`
	InstanceId           *string   `json:"instanceId,omitempty"`
	CacheNodeIds         []*string `json:"cacheNodeIds,omitempty"`
	ForceRemoveOnOffline *bool     `json:"forceRemoveOnOffline,omitempty"`
}
