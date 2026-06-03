package rapidfs

type CancelMetaSyncJobRequest struct {
	ClientToken    *string `json:"-"`
	InstanceId     *string `json:"instanceId,omitempty"`
	DataSrcId      *string `json:"dataSrcId,omitempty"`
	MetaSyncRuleId *string `json:"metaSyncRuleId,omitempty"`
}
