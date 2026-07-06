package aihc

type JobResult struct {
	JobId    *string `json:"jobId,omitempty"`
	Success  *bool   `json:"success,omitempty"`
	ErrorMsg *string `json:"errorMsg,omitempty"`
}
