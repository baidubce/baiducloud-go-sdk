package ocr

type CorrectEduResult struct {
	TaskId        *string        `json:"task_id,omitempty"`
	IsAllFinished *bool          `json:"isAllFinished,omitempty"`
	Status        *string        `json:"status,omitempty"`
	StatResult    *StatResult    `json:"stat_result,omitempty"`
	ImageResults  []*ImageResult `json:"imageResults,omitempty"`
}
