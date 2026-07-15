package cloudassistant

type ActionRunListRequest struct {
	Locale    *string       `json:"-"`
	PageNo    *int32        `json:"pageNo,omitempty"`
	PageSize  *int32        `json:"pageSize,omitempty"`
	Sort      *string       `json:"sort,omitempty"`
	Ascending *bool         `json:"ascending,omitempty"`
	Action    *ActionFilter `json:"action,omitempty"`
	State     *string       `json:"state,omitempty"`
	RunId     *string       `json:"runId,omitempty"`
	StartTime *int32        `json:"startTime,omitempty"`
	EndTime   *int32        `json:"endTime,omitempty"`
}
