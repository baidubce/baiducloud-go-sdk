package rapidfs

type DescribeMetaSyncJobsRequest struct {
	Action         *string `json:"-"`
	InstanceId     *string `json:"instanceId,omitempty"`
	DataSrcId      *string `json:"dataSrcId,omitempty"`
	MetaSyncRuleId *string `json:"metaSyncRuleId,omitempty"`
	MaxKeys        *int32  `json:"maxKeys,omitempty"`
	Marker         *string `json:"marker,omitempty"`
}
