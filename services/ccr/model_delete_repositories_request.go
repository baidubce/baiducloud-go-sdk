package ccr

type DeleteRepositoriesRequest struct {
	InstanceId  *string   `json:"-"`
	ProjectName *string   `json:"-"`
	Items       []*string `json:"items,omitempty"`
}
