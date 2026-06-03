package rapidfs

type MetaSyncRuleInfo struct {
	MetaSyncRuleId   *string `json:"metaSyncRuleId,omitempty"`
	MetaSyncRuleName *string `json:"metaSyncRuleName,omitempty"`
	InstanceId       *string `json:"instanceId,omitempty"`
	InstanceName     *string `json:"instanceName,omitempty"`
	DataSrcName      *string `json:"dataSrcName,omitempty"`
	DataSrcId        *string `json:"dataSrcId,omitempty"`
	RapidfsType      *string `json:"type,omitempty"`
	Directory        *string `json:"directory,omitempty"`
	IntervalMinutes  *int32  `json:"IntervalMinutes,omitempty"`
	Status           *string `json:"status,omitempty"`
	CreateTime       *string `json:"createTime,omitempty"`
	Description      *string `json:"description,omitempty"`
	LastJobStatus    *string `json:"lastJobStatus,omitempty"`
	LastJobStartTime *string `json:"lastJobStartTime,omitempty"`
	LastJobEndTime   *string `json:"lastJobEndTime,omitempty"`
}
