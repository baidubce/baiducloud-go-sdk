package ccr

type TaskResult struct {
	DestResource *string `json:"destResource,omitempty"`
	EndTime      *string `json:"endTime,omitempty"`
	ExecutionId  *int32  `json:"executionId,omitempty"`
	Id           *int32  `json:"id,omitempty"`
	JobId        *string `json:"jobId,omitempty"`
	Operation    *string `json:"operation,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	SrcResource  *string `json:"srcResource,omitempty"`
	StartTime    *string `json:"startTime,omitempty"`
	Status       *string `json:"status,omitempty"`
}
