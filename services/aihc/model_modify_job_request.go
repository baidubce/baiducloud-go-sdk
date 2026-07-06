package aihc

type ModifyJobRequest struct {
	ResourcePoolId *string `json:"-"`
	QueueID        *string `json:"-"`
	JobId          *string `json:"jobId,omitempty"`
	Priority       *string `json:"priority,omitempty"`
}
