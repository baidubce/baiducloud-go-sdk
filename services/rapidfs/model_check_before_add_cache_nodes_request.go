package rapidfs

type CheckBeforeAddCacheNodesRequest struct {
	InstanceId  *string             `json:"instanceId,omitempty"`
	RapidfsType *string             `json:"type,omitempty"`
	CacheNodes  []*AddCacheNodeInfo `json:"cacheNodes,omitempty"`
}
