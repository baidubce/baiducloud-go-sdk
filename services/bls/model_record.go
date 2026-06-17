package bls

type Record struct {
	StartTime     *string `json:"startTime,omitempty"`
	EndTime       *string `json:"endTime,omitempty"`
	FinishedCount *int32  `json:"finishedCount,omitempty"`
}
