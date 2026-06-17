package bls

type ListAlarmExecutionStatsRequest struct {
	PolicyId      *string   `json:"policyId,omitempty"`
	PolicyName    *string   `json:"policyName,omitempty"`
	LogStoreName  *string   `json:"logStoreName,omitempty"`
	States        []*string `json:"states,omitempty"`
	StartDateTime *string   `json:"startDateTime,omitempty"`
	EndDateTime   *string   `json:"endDateTime,omitempty"`
	OrderBy       *string   `json:"orderBy,omitempty"`
	Order         *string   `json:"order,omitempty"`
	PageNo        *int32    `json:"pageNo,omitempty"`
	PageSize      *int32    `json:"pageSize,omitempty"`
}
