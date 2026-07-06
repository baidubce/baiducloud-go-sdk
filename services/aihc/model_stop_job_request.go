package aihc

type StopJobRequest struct {
	ResourcePoolId *string `json:"-"`
	QueueID        *string `json:"-"`
	JobId          *string `json:"jobId,omitempty"`
}
