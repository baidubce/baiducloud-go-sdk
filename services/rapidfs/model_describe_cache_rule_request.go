package rapidfs

type DescribeCacheRuleRequest struct {
	InstanceId  *string `json:"instanceId,omitempty"`
	DataSrcId   *string `json:"dataSrcId,omitempty"`
	CacheRuleId *string `json:"cacheRuleId,omitempty"`
}
