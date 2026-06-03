package ccr

type ReExecuteTriggerTaskRequest struct {
	InstanceId *string `json:"-"`
	PolicyId   *string `json:"-"`
	JobId      *string `json:"-"`
}
