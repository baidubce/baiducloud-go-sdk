package ccr

type DeleteChartsRequest struct {
	InstanceId  *string   `json:"-"`
	ProjectName *string   `json:"-"`
	Items       []*string `json:"items,omitempty"`
}
