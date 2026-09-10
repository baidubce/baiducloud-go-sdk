package cce

type Task struct {
	Id          *string        `json:"id,omitempty"`
	CceType     *string        `json:"type,omitempty"`
	Description *string        `json:"description,omitempty"`
	StartTime   *string        `json:"startTime,omitempty"`
	FinishTime  *string        `json:"finishTime,omitempty"`
	Phase       *string        `json:"phase,omitempty"`
	Processes   []*interface{} `json:"processes,omitempty"`
}
