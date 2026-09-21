package dbsc

type BigKeyTask struct {
	NodeId     *string `json:"nodeId,omitempty"`
	TaskId     *int32  `json:"taskId,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	EndTime    *string `json:"endTime,omitempty"`
	StartTime  *string `json:"startTime,omitempty"`
	Status     *int32  `json:"status,omitempty"`
	StatusDesc *string `json:"statusDesc,omitempty"`
}
