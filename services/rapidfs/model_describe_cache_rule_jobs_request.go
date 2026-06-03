package rapidfs

type DescribeCacheRuleJobsRequest struct {
	InstanceId  *string `json:"instanceId,omitempty"`
	DataSrcId   *string `json:"dataSrcId,omitempty"`
	CacheRuleId *string `json:"cacheRuleId,omitempty"`
	MaxKeys     *int32  `json:"maxKeys,omitempty"`
	Marker      *string `json:"marker,omitempty"`
}
