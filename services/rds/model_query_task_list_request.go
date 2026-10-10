package rds

type QueryTaskListRequest struct {
	PageSize     *string `json:"pageSize,omitempty"`
	PageNo       *string `json:"pageNo,omitempty"`
	InstanceId   *string `json:"instanceId,omitempty"`
	InstanceName *string `json:"instanceName,omitempty"`
	TaskId       *int32  `json:"taskId,omitempty"`
	TaskType     *string `json:"taskType,omitempty"`
	TaskStatus   *string `json:"taskStatus,omitempty"`
	StartTime    *string `json:"startTime,omitempty"`
	EndTime      *string `json:"endTime,omitempty"`
}
