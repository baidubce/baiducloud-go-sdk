package apm

type ApmDescribeAlarmPoliciesRequest struct {
	PolicyName *string `json:"policyName,omitempty"`
	PolicyId   *string `json:"policyId,omitempty"`
	State      *string `json:"state,omitempty"`
	Level      *string `json:"level,omitempty"`
	MetricKind *string `json:"metricKind,omitempty"`
	OrderBy    *string `json:"orderBy,omitempty"`
	Order      *string `json:"order,omitempty"`
	PageNo     *int32  `json:"pageNo,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty"`
}
