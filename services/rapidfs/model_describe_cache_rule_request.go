package rapidfs

type DescribeCacheRuleRequest struct {
	Action      *string `json:"-"`
	InstanceId  *string `json:"instanceId,omitempty"`
	DataSrcId   *string `json:"dataSrcId,omitempty"`
	CacheRuleId *string `json:"cacheRuleId,omitempty"`
}
