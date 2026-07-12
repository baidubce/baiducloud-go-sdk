package oos

type Execution struct {
	Id                *string                 `json:"id,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	Template          *Template               `json:"template,omitempty"`
	TemplateDeleted   *bool                   `json:"templateDeleted,omitempty"`
	Parallelism       *int32                  `json:"parallelism,omitempty"`
	Manually          *bool                   `json:"manually,omitempty"`
	CreatedTimestamp  *int64                  `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp  *int64                  `json:"updatedTimestamp,omitempty"`
	FinishedTimestamp *int64                  `json:"finishedTimestamp,omitempty"`
	State             *string                 `json:"state,omitempty"`
	Properties        *interface{}            `json:"properties,omitempty"`
	Tasks             []*ExecutionTaskSummary `json:"tasks,omitempty"`
	Tags              []*Tag                  `json:"tags,omitempty"`
	Trigger           *string                 `json:"trigger,omitempty"`
	Reason            *string                 `json:"reason,omitempty"`
	ErrorCode         *string                 `json:"errorCode,omitempty"`
	EventExecution    *EventExecution         `json:"eventExecution,omitempty"`
	CronExecution     *CronExecution          `json:"cronExecution,omitempty"`
	Locale            *string                 `json:"locale,omitempty"`
}
