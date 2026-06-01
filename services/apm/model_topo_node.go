package apm

type TopoNode struct {
	ServiceName       *string  `json:"serviceName,omitempty"`
	Language          *string  `json:"language,omitempty"`
	ApmType           *string  `json:"type,omitempty"`
	Inferred          *bool    `json:"inferred,omitempty"`
	Component         *string  `json:"component,omitempty"`
	Requests          *int32   `json:"requests,omitempty"`
	RequestsPerSecond *float32 `json:"requestsPerSecond,omitempty"`
	Errors            *int32   `json:"errors,omitempty"`
	ErrorRate         *float32 `json:"errorRate,omitempty"`
	DurationSeconds   *float32 `json:"durationSeconds,omitempty"`
	State             *string  `json:"state,omitempty"`
}
