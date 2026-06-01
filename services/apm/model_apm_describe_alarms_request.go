package apm

type ApmDescribeAlarmsRequest struct {
	BeginDatetime *string `json:"beginDatetime,omitempty"`
	EndDatetime   *string `json:"endDatetime,omitempty"`
	PolicyName    *string `json:"policyName,omitempty"`
	Level         *string `json:"level,omitempty"`
	MetricKind    *string `json:"metricKind,omitempty"`
	State         *string `json:"state,omitempty"`
	OrderBy       *string `json:"orderBy,omitempty"`
	Order         *string `json:"order,omitempty"`
	PageNo        *int32  `json:"pageNo,omitempty"`
	PageSize      *int32  `json:"pageSize,omitempty"`
}
