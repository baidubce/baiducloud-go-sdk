package rapidfs

type DescribeCacheNodeRequest struct {
	Action      *string `json:"-"`
	InstanceId  *string `json:"instanceId,omitempty"`
	CacheNodeId *string `json:"cacheNodeId,omitempty"`
}
