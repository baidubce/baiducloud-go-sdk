package oos

type CronExecution struct {
	Namespace             *string      `json:"namespace,omitempty"`
	Description           *string      `json:"description,omitempty"`
	Name                  *string      `json:"name,omitempty"`
	Template              *Template    `json:"template,omitempty"`
	TemplateDeleted       *bool        `json:"templateDeleted,omitempty"`
	Properties            *interface{} `json:"properties,omitempty"`
	Tags                  []*Tag       `json:"tags,omitempty"`
	Cron                  *string      `json:"cron,omitempty"`
	Period                *Period      `json:"period,omitempty"`
	DependOnPast          *bool        `json:"dependOnPast,omitempty"`
	ScheduleTimeout       *int64       `json:"scheduleTimeout,omitempty"`
	CreatedTimestamp      *int64       `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp      *int64       `json:"updatedTimestamp,omitempty"`
	NextScheduleTimestamp *int64       `json:"nextScheduleTimestamp,omitempty"`
	BeginTimestamp        *int64       `json:"beginTimestamp,omitempty"`
	EndTimestamp          *int64       `json:"endTimestamp,omitempty"`
	State                 *string      `json:"state,omitempty"`
	Locale                *string      `json:"locale,omitempty"`
}
