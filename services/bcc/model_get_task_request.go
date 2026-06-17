package bcc

type GetTaskRequest struct {
	TaskIds []*string `json:"taskIds,omitempty"`
	MaxKeys *int32    `json:"maxKeys,omitempty"`
}
