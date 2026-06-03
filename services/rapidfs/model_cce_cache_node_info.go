package rapidfs

type CCECacheNodeInfo struct {
	CacheDeployGroup *string `json:"cacheDeployGroup,omitempty"`
	CceClusterId     *string `json:"cceClusterId,omitempty"`
	HostBccId        *string `json:"hostBccId,omitempty"`
	HostBccName      *string `json:"hostBccName,omitempty"`
	HostZone         *string `json:"hostZone,omitempty"`
}
