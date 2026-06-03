package ccr

type ListTriggerTasksRequest struct {
	InstanceId *string `json:"-"`
	PolicyId   *string `json:"-"`
	PageNo     *int32  `json:"-"`
	PageSize   *int32  `json:"-"`
}
