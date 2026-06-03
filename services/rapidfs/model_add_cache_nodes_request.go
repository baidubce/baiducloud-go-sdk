package rapidfs

type AddCacheNodesRequest struct {
	ClientToken *string             `json:"-"`
	InstanceId  *string             `json:"instanceId,omitempty"`
	RapidfsType *string             `json:"type,omitempty"`
	CacheNodes  []*AddCacheNodeInfo `json:"cacheNodes,omitempty"`
}
