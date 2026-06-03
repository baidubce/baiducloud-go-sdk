package rapidfs

type CreateMetaSyncRuleRequest struct {
	Action           *string `json:"-"`
	ClientToken      *string `json:"-"`
	InstanceId       *string `json:"instanceId,omitempty"`
	DataSrcId        *string `json:"dataSrcId,omitempty"`
	MetaSyncRuleName *string `json:"metaSyncRuleName,omitempty"`
	RapidfsType      *string `json:"type,omitempty"`
	Directory        *string `json:"directory,omitempty"`
	IntervalMinutes  *int32  `json:"intervalMinutes,omitempty"`
	ExecuteOnCreate  *bool   `json:"executeOnCreate,omitempty"`
	EnableOnCreate   *bool   `json:"enableOnCreate,omitempty"`
	Description      *string `json:"description,omitempty"`
}
