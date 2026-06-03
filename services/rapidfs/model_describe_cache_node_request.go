package rapidfs

type DescribeCacheNodeRequest struct {
	InstanceId  *string `json:"instanceId,omitempty"`
	CacheNodeId *string `json:"cacheNodeId,omitempty"`
}
