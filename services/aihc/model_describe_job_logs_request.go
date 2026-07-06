package aihc

type DescribeJobLogsRequest struct {
	ResourcePoolId *string `json:"-"`
	QueueID        *string `json:"-"`
	JobId          *string `json:"jobId,omitempty"`
	PodName        *string `json:"podName,omitempty"`
	Keywords       *string `json:"keywords,omitempty"`
	StartTime      *string `json:"startTime,omitempty"`
	EndTime        *string `json:"endTime,omitempty"`
	MaxLines       *string `json:"maxLines,omitempty"`
	ChunkSize      *string `json:"chunkSize,omitempty"`
	Marker         *string `json:"marker,omitempty"`
}
