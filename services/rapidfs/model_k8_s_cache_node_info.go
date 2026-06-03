package rapidfs

type K8SCacheNodeInfo struct {
	CacheDeployGroup *string `json:"cacheDeployGroup,omitempty"`
	K8sControllerId  *string `json:"k8sControllerId,omitempty"`
}
