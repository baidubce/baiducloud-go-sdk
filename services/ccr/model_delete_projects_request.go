package ccr

type DeleteProjectsRequest struct {
	InstanceId *string   `json:"-"`
	Items      []*string `json:"items,omitempty"`
}
