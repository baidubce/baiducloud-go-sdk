package ccr

type StopInstanceSyncRequest struct {
	InstanceId  *string `json:"-"`
	ExecutionId *string `json:"-"`
}
