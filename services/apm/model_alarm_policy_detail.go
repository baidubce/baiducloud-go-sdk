package apm

type AlarmPolicyDetail struct {
	Id                                   *string        `json:"id,omitempty"`
	Name                                 *string        `json:"name,omitempty"`
	Content                              *string        `json:"content,omitempty"`
	State                                *string        `json:"state,omitempty"`
	CreatedTimestamp                     *string        `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp                     *string        `json:"updatedTimestamp,omitempty"`
	Target                               *AlarmTarget   `json:"target,omitempty"`
	MetricKind                           *string        `json:"metricKind,omitempty"`
	Rule                                 *AlarmRule     `json:"rule,omitempty"`
	RuleContent                          *string        `json:"ruleContent,omitempty"`
	Region                               *string        `json:"region,omitempty"`
	Filters                              []*AlarmFilter `json:"filters,omitempty"`
	PendingCount                         *int32         `json:"pendingCount,omitempty"`
	RenotifyIntervalInMinutes            *int32         `json:"renotifyIntervalInMinutes,omitempty"`
	RenotifyCount                        *int32         `json:"renotifyCount,omitempty"`
	NotifyRecovery                       *bool          `json:"notifyRecovery,omitempty"`
	OnMissingData                        *string        `json:"onMissingData,omitempty"`
	NoDataNotifyPendingIntervalInMinutes *int32         `json:"noDataNotifyPendingIntervalInMinutes,omitempty"`
	Level                                *string        `json:"level,omitempty"`
	Actions                              []*AlarmAction `json:"actions,omitempty"`
}
