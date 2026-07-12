package oos

type GetExecutionListV2Request struct {
	Locale             *string      `json:"-"`
	Namespace          *string      `json:"namespace,omitempty"`
	Template           *interface{} `json:"template,omitempty"`
	State              *string      `json:"state,omitempty"`
	Trigger            *string      `json:"trigger,omitempty"`
	CronExecutionName  *string      `json:"cronExecutionName,omitempty"`
	EventExecutionName *string      `json:"eventExecutionName,omitempty"`
	StartTime          *int32       `json:"startTime,omitempty"`
	EndTime            *int32       `json:"endTime,omitempty"`
	Sort               *string      `json:"sort,omitempty"`
	Ascending          *bool        `json:"ascending,omitempty"`
	PageNo             *int32       `json:"pageNo,omitempty"`
	PageSize           *int32       `json:"pageSize,omitempty"`
}
