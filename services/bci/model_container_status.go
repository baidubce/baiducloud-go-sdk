package bci

type ContainerStatus struct {
	State        *string `json:"state,omitempty"`
	Reason       *string `json:"reason,omitempty"`
	Message      *string `json:"message,omitempty"`
	StartTime    *string `json:"startTime,omitempty"`
	FinishTime   *string `json:"finishTime,omitempty"`
	DetailStatus *string `json:"detailStatus,omitempty"`
	ExitCode     *int32  `json:"exitCode,omitempty"`
}
