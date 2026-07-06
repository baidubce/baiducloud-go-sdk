package aihc

type BatchStopTrainingTasksV2Request struct {
	QueueID        *string                   `json:"-"`
	ResourcePoolId *string                   `json:"-"`
	JobList        []*map[string]interface{} `json:"jobList,omitempty"`
	JobListJobId   *string                   `json:"jobList[].jobId,omitempty"`
}
