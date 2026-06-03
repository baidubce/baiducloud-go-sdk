package rapidfs

type AIHCCacheNodeInfo struct {
	CacheDeployGroup   *string `json:"cacheDeployGroup,omitempty"`
	AihcResourcePoolId *string `json:"aihcResourcePoolId,omitempty"`
	HostNodeId         *string `json:"hostNodeId,omitempty"`
	HostNodeName       *string `json:"hostNodeName,omitempty"`
	HostZone           *string `json:"hostZone,omitempty"`
}
