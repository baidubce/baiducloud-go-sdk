package ocr

type HandwritingCompositionCreateTaskResult struct {
	TaskId        *string `json:"task_id,omitempty"`
	WaitTaskCount *int32  `json:"wait_task_count,omitempty"`
}
