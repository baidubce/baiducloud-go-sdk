package cloudassistant

type LogResult struct {
	Logs       []*Log  `json:"logs,omitempty"`
	NextCursor *int32  `json:"nextCursor,omitempty"`
	State      *string `json:"state,omitempty"`
	ChildId    *string `json:"childId,omitempty"`
}
