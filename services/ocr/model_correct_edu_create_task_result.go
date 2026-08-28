package ocr

type CorrectEduCreateTaskResult struct {
	TaskId        *string `json:"task_id,omitempty"`
	WaitTaskCount *int32  `json:"wait_task_count,omitempty"`
}
