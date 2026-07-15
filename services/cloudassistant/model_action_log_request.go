package cloudassistant

type ActionLogRequest struct {
	RunId   *string `json:"runId,omitempty"`
	ChildId *string `json:"childId,omitempty"`
	Cursor  *int32  `json:"cursor,omitempty"`
}
