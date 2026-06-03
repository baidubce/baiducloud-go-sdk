package rapidfs

type DescribeCacheDeployGroupRequest struct {
	Action               *string `json:"-"`
	InstanceId           *string `json:"instanceId,omitempty"`
	CacheDeployGroupName *string `json:"cacheDeployGroupName,omitempty"`
	CacheDeployGroupNS   *string `json:"cacheDeployGroupNS,omitempty"`
}
