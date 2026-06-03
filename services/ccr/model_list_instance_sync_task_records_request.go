package ccr

type ListInstanceSyncTaskRecordsRequest struct {
	InstanceId  *string `json:"-"`
	ExecutionId *string `json:"-"`
	PageNo      *int32  `json:"-"`
	PageSize    *int32  `json:"-"`
}
