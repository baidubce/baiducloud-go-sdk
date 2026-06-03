package rapidfs

type DescribeAllocatableDataSrcQuotaRequest struct {
	Action     *string `json:"-"`
	InstanceId *string `json:"instanceId,omitempty"`
	DataSrcId  *string `json:"dataSrcId,omitempty"`
}
