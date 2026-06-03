package ccr

type DeleteChartVersionsRequest struct {
	InstanceId  *string   `json:"-"`
	ProjectName *string   `json:"-"`
	ChartName   *string   `json:"-"`
	Items       []*string `json:"items,omitempty"`
}
