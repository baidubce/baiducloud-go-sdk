package ccr

type ListImageMigrationTaskRecordsRequest struct {
	InstanceId  *string `json:"-"`
	ExecutionId *string `json:"-"`
	PageNo      *int32  `json:"-"`
	PageSize    *int32  `json:"-"`
}
