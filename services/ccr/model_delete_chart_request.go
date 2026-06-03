package ccr

type DeleteChartRequest struct {
	InstanceId  *string `json:"-"`
	ProjectName *string `json:"-"`
	ChartName   *string `json:"-"`
}
