package ocr

type HandwritingGetResult struct {
	TaskId       *string                    `json:"task_id,omitempty"`
	Status       *string                    `json:"status,omitempty"`
	CreatedTime  *int32                     `json:"created_time,omitempty"`
	StartedTime  *int32                     `json:"started_time,omitempty"`
	FinishedTime *int32                     `json:"finished_time,omitempty"`
	Duration     *int32                     `json:"duration,omitempty"`
	Result       *HandwritingGetEssayResult `json:"result,omitempty"`
}
