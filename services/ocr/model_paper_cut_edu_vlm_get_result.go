package ocr

type PaperCutEduVlmGetResult struct {
	TaskId     *string                       `json:"task_id,omitempty"`
	Status     *string                       `json:"status,omitempty"`
	QusResults []*PaperCutEduVlmGetQusResult `json:"qus_results,omitempty"`
}
