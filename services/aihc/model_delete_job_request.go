package aihc

type DeleteJobRequest struct {
	ResourcePoolId *string `json:"-"`
	QueueID        *string `json:"-"`
	JobId          *string `json:"jobId,omitempty"`
}
