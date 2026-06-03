package rapidfs

type DescribeAllocatableDataSrcQuotaRequest struct {
	InstanceId *string `json:"instanceId,omitempty"`
	DataSrcId  *string `json:"dataSrcId,omitempty"`
}
