package rapidfs

type RestartCacheNodesRequest struct {
	Action       *string   `json:"-"`
	ClientToken  *string   `json:"-"`
	InstanceId   *string   `json:"instanceId,omitempty"`
	CacheNodeIds []*string `json:"cacheNodeIds,omitempty"`
}
