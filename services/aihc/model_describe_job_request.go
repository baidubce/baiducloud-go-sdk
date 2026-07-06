package aihc

type DescribeJobRequest struct {
	ResourcePoolId *string `json:"-"`
	QueueID        *string `json:"-"`
	JobId          *string `json:"jobId,omitempty"`
	NeedDetail     *bool   `json:"needDetail,omitempty"`
}
