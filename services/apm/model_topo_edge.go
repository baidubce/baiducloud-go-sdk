package apm

type TopoEdge struct {
	Source            *string  `json:"source,omitempty"`
	Target            *string  `json:"target,omitempty"`
	ApmType           *string  `json:"type,omitempty"`
	Requests          *int32   `json:"requests,omitempty"`
	RequestsPerSecond *float32 `json:"requestsPerSecond,omitempty"`
	Errors            *int32   `json:"errors,omitempty"`
	ErrorRate         *float32 `json:"errorRate,omitempty"`
	DurationSeconds   *float32 `json:"durationSeconds,omitempty"`
	State             *string  `json:"state,omitempty"`
}
