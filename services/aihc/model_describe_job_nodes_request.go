package aihc

type DescribeJobNodesRequest struct {
	ResourcePoolId *string `json:"-"`
	QueueID        *string `json:"-"`
	JobId          *string `json:"jobId,omitempty"`
}
