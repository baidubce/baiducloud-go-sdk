package rapidfs

type AddCacheNodeInfo struct {
	Ip         *string           `json:"ip,omitempty"`
	Config     *string           `json:"config,omitempty"`
	DiskInfos  []*DiskInfo       `json:"diskInfos,omitempty"`
	DeployPath *string           `json:"deployPath,omitempty"`
	BccInfo    *BCCCacheNodeInfo `json:"bccInfo,omitempty"`
	IdcInfo    *IDCCacheNodeInfo `json:"idcInfo,omitempty"`
}
