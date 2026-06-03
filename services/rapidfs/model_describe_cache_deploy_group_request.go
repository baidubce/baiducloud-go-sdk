package rapidfs

type DescribeCacheDeployGroupRequest struct {
	InstanceId           *string `json:"instanceId,omitempty"`
	CacheDeployGroupName *string `json:"cacheDeployGroupName,omitempty"`
	CacheDeployGroupNS   *string `json:"cacheDeployGroupNS,omitempty"`
}
