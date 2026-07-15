package cloudassistant

type ActionRun struct {
	Id                *string     `json:"id,omitempty"`
	State             *string     `json:"state,omitempty"`
	Action            *Action     `json:"action,omitempty"`
	CreatedTimestamp  *int64      `json:"createdTimestamp,omitempty"`
	FinishedTimestamp *int64      `json:"finishedTimestamp,omitempty"`
	Statistics        *Statistics `json:"statistics,omitempty"`
	Children          []*ChildRun `json:"children,omitempty"`
	TotalCount        *int32      `json:"totalCount,omitempty"`
}
