package aihc

type JobItem struct {
	Jobid                 *string        `json:"jobid,omitempty"`
	UserId                *string        `json:"userId,omitempty"`
	Name                  *string        `json:"name,omitempty"`
	Status                *string        `json:"status,omitempty"`
	JobType               *string        `json:"jobType,omitempty"`
	ResourcePoolId        *string        `json:"resourcePoolId,omitempty"`
	Queue                 *string        `json:"queue,omitempty"`
	Job                   *JobSpec       `json:"Job,omitempty"`
	CreatedAt             *string        `json:"createdAt,omitempty"`
	FinishedAt            *string        `json:"finishedAt,omitempty"`
	Datasources           []*DataSource  `json:"datasources,omitempty"`
	Labels                []*Label       `json:"labels,omitempty"`
	Priority              *string        `json:"priority,omitempty"`
	EnableBccl            *bool          `json:"enableBccl,omitempty"`
	EnableBcclStatus      *string        `json:"enableBcclStatus,omitempty"`
	EnableBcclErrorReason *string        `json:"enableBcclErrorReason,omitempty"`
	EnableFaultTolerance  *bool          `json:"enableFaultTolerance,omitempty"`
	FaultToleranceArgs    *string        `json:"faultToleranceArgs,omitempty"`
	Pods                  []*Pod         `json:"pods,omitempty"`
	HistoryPods           []*Pod         `json:"historyPods,omitempty"`
	JobTimeLine           []*JobTimeLine `json:"jobTimeLine,omitempty"`
}
