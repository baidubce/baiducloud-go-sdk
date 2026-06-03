package rapidfs

type DataSrcQuotaInfo struct {
	InstanceId          *string `json:"instanceId,omitempty"`
	AllocatableQuotaMiB *int64  `json:"allocatableQuotaMiB,omitempty"`
}
