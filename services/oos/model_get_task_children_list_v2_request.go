package oos

type GetTaskChildrenListV2Request struct {
	Locale      *string   `json:"-"`
	ExecutionId *string   `json:"executionId,omitempty"`
	TaskId      *string   `json:"taskId,omitempty"`
	Namespace   *string   `json:"namespace,omitempty"`
	States      []*string `json:"states,omitempty"`
	PageNo      *int32    `json:"pageNo,omitempty"`
	PageSize    *int32    `json:"pageSize,omitempty"`
}
