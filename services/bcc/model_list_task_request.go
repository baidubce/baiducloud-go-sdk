package bcc

type ListTaskRequest struct {
	TaskIds     []*string `json:"taskIds,omitempty"`
	Marker      *string   `json:"marker,omitempty"`
	MaxKeys     *int32    `json:"maxKeys,omitempty"`
	TaskAction  *string   `json:"taskAction,omitempty"`
	TaskStatus  *string   `json:"taskStatus,omitempty"`
	StartTime   *string   `json:"startTime,omitempty"`
	EndTime     *string   `json:"endTime,omitempty"`
	ResourceIds []*string `json:"resourceIds,omitempty"`
}
