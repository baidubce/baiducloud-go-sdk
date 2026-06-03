package rapidfs

type CacheRuleInfo struct {
	CacheRuleId      *string `json:"cacheRuleId,omitempty"`
	CacheRuleName    *string `json:"cacheRuleName,omitempty"`
	InstanceId       *string `json:"instanceId,omitempty"`
	InstanceName     *string `json:"instanceName,omitempty"`
	DataSrcId        *string `json:"dataSrcId,omitempty"`
	DataSrcName      *string `json:"dataSrcName,omitempty"`
	RapidfsType      *string `json:"type,omitempty"`
	Directory        *string `json:"directory,omitempty"`
	Status           *string `json:"status,omitempty"`
	CreateTime       *string `json:"createTime,omitempty"`
	Description      *string `json:"description,omitempty"`
	LastJobStatus    *string `json:"lastJobStatus,omitempty"`
	LastJobSize      *int64  `json:"lastJobSize,omitempty"`
	LastJobStartTime *string `json:"lastJobStartTime,omitempty"`
	LastJobEndTime   *string `json:"lastJobEndTime,omitempty"`
}
