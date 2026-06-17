package bls

type ListAlarmExecutionsRequest struct {
	PolicyId      *string   `json:"policyId,omitempty"`
	LogStoreName  *string   `json:"logStoreName,omitempty"`
	State         *string   `json:"state,omitempty"`
	NoticeStates  []*string `json:"noticeStates,omitempty"`
	StartDateTime *string   `json:"startDateTime,omitempty"`
	EndDateTime   *string   `json:"endDateTime,omitempty"`
	OrderBy       *string   `json:"orderBy,omitempty"`
	Order         *string   `json:"order,omitempty"`
	PageNo        *int32    `json:"pageNo,omitempty"`
	PageSize      *int32    `json:"pageSize,omitempty"`
}
