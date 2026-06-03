package ccr

type GetInstanceSyncTaskLogsRequest struct {
	InstanceId  *string `json:"-"`
	ExecutionId *string `json:"-"`
	TaskId      *string `json:"-"`
}
