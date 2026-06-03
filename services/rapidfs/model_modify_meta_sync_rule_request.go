package rapidfs

type ModifyMetaSyncRuleRequest struct {
	Action           *string `json:"-"`
	ClientToken      *string `json:"-"`
	InstanceId       *string `json:"instanceId,omitempty"`
	DataSrcId        *string `json:"dataSrcId,omitempty"`
	MetaSyncRuleId   *string `json:"metaSyncRuleId,omitempty"`
	MetaSyncRuleName *string `json:"metaSyncRuleName,omitempty"`
	IntervalMinutes  *int32  `json:"intervalMinutes,omitempty"`
	Description      *string `json:"description,omitempty"`
}
