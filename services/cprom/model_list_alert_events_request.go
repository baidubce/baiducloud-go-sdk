package cprom

type ListAlertEventsRequest struct {
	StartTime         *int32  `json:"-"`
	EndTime           *int32  `json:"-"`
	PageNo            *int32  `json:"-"`
	PageSize          *int32  `json:"-"`
	MonitorInstanceId *string `json:"-"`
	AlertingRuleId    *string `json:"-"`
	AlertingRuleName  *string `json:"-"`
	NotifyRuleId      *string `json:"-"`
	NotifyRuleName    *string `json:"-"`
	Severity          *string `json:"-"`
	Status            *string `json:"-"`
	Expr              *string `json:"-"`
	OrderBy           *string `json:"-"`
	Order             *string `json:"-"`
	AlarmTags         *string `json:"-"`
}
