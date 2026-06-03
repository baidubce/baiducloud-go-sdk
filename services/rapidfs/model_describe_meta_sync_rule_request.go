package rapidfs

type DescribeMetaSyncRuleRequest struct {
	Action         *string `json:"-"`
	InstanceId     *string `json:"instanceId,omitempty"`
	DataSrcId      *string `json:"dataSrcId,omitempty"`
	MetaSyncRuleId *string `json:"metaSyncRuleId,omitempty"`
}
