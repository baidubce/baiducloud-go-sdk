package ccr

type GetImageMigrationTaskLogsRequest struct {
	InstanceId  *string `json:"-"`
	ExecutionId *string `json:"-"`
	TaskId      *string `json:"-"`
}
