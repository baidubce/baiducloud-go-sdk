package rapidfs

type RestartCacheNodesRequest struct {
	ClientToken  *string   `json:"-"`
	InstanceId   *string   `json:"instanceId,omitempty"`
	CacheNodeIds []*string `json:"cacheNodeIds,omitempty"`
}
