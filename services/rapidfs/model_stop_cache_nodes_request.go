package rapidfs

type StopCacheNodesRequest struct {
	ClientToken           *string   `json:"-"`
	InstanceId            *string   `json:"instanceId,omitempty"`
	CacheNodeIds          []*string `json:"cacheNodeIds,omitempty"`
	MigrateDataBeforeStop *bool     `json:"migrateDataBeforeStop,omitempty"`
}
