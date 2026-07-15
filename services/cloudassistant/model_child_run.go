package cloudassistant

type ChildRun struct {
	Id                *string      `json:"id,omitempty"`
	State             *string      `json:"state,omitempty"`
	Target            *interface{} `json:"target,omitempty"`
	CreatedTimestamp  *int64       `json:"createdTimestamp,omitempty"`
	FinishedTimestamp *int64       `json:"finishedTimestamp,omitempty"`
	Output            *Output      `json:"output,omitempty"`
	Log               *string      `json:"log,omitempty"`
	ErrorCode         *string      `json:"errorCode,omitempty"`
	FailReason        *string      `json:"failReason,omitempty"`
}
