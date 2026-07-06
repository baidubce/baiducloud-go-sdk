package aihc

type StopBatchResult struct {
	SuccessNum  *int32       `json:"successNum,omitempty"`
	FailedNum   *int32       `json:"failedNum,omitempty"`
	Success     *bool        `json:"success,omitempty"`
	SuccessList []*JobResult `json:"successList,omitempty"`
	FailedList  []*JobResult `json:"failedList,omitempty"`
}
