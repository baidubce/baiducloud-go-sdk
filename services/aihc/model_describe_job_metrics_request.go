package aihc

type DescribeJobMetricsRequest struct {
	ResourcePoolId *string `json:"-"`
	QueueID        *string `json:"-"`
	JobId          *string `json:"jobId,omitempty"`
	StartTime      *string `json:"startTime,omitempty"`
	EndTime        *string `json:"endTime,omitempty"`
	TimeStep       *string `json:"timeStep,omitempty"`
	MetricType     *string `json:"metricType,omitempty"`
	RateInterval   *string `json:"rateInterval,omitempty"`
}
