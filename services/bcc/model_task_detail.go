package bcc

type TaskDetail struct {
	TaskId               *string                 `json:"taskId,omitempty"`
	TaskAction           *string                 `json:"taskAction,omitempty"`
	TaskStatus           *string                 `json:"taskStatus,omitempty"`
	CreatedTime          *string                 `json:"createdTime,omitempty"`
	FinishedTime         *string                 `json:"finishedTime,omitempty"`
	TotalCount           *int32                  `json:"totalCount,omitempty"`
	SuccessCount         *int32                  `json:"successCount,omitempty"`
	FailedCount          *int32                  `json:"failedCount,omitempty"`
	OperationProgressSet []*OperationProgressSet `json:"operationProgressSet,omitempty"`
}
