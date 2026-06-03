package rapidfs

type StartCacheNodesRequest struct {
	ClientToken  *string   `json:"-"`
	InstanceId   *string   `json:"instanceId,omitempty"`
	CacheNodeIds []*string `json:"cacheNodeIds,omitempty"`
}
