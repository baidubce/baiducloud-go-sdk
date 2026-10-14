package rds

type Task struct {
	TaskId       *int32      `json:"taskId,omitempty"`
	TaskType     *string     `json:"taskType,omitempty"`
	TaskName     *string     `json:"taskName,omitempty"`
	InstanceId   *string     `json:"instanceId,omitempty"`
	InstanceName *string     `json:"instanceName,omitempty"`
	UserId       *string     `json:"userId,omitempty"`
	Region       *string     `json:"region,omitempty"`
	TaskStatus   *string     `json:"taskStatus,omitempty"`
	CreateTime   *string     `json:"createTime,omitempty"`
	UpdateTime   *string     `json:"updateTime,omitempty"`
	FinishTime   *string     `json:"finishTime,omitempty"`
	CancelFlag   *int32      `json:"cancelFlag,omitempty"`
	Progress     []*Progress `json:"progress,omitempty"`
}
