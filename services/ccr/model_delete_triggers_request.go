package ccr

type DeleteTriggersRequest struct {
	InstanceId *string  `json:"-"`
	Items      []*int32 `json:"items,omitempty"`
}
