package aihc

type CreateJobRequest struct {
	ResourcePoolId     *string            `json:"-"`
	QueueID            *string            `json:"-"`
	Name               *string            `json:"name,omitempty"`
	Queue              *string            `json:"queue,omitempty"`
	JobType            *string            `json:"jobType,omitempty"`
	JobSpec            *JobSpec           `json:"jobSpec,omitempty"`
	Command            *string            `json:"command,omitempty"`
	Labels             []*Label           `json:"labels,omitempty"`
	Priority           *string            `json:"priority,omitempty"`
	Datasources        []*DataSource      `json:"datasources,omitempty"`
	EnableBccl         *bool              `json:"enableBccl,omitempty"`
	FaultTolerance     *bool              `json:"faultTolerance,omitempty"`
	FaultToleranceArgs *string            `json:"faultToleranceArgs,omitempty"`
	TensorboardConfig  *TensorboardConfig `json:"tensorboardConfig,omitempty"`
	AlertConfig        *AlertConfig       `json:"alertConfig,omitempty"`
	RetentionPeriod    *string            `json:"retentionPeriod,omitempty"`
	AdvancedSettings   *AdvancedSettings  `json:"advancedSettings,omitempty"`
}
