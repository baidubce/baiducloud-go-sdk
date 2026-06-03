package ccr

type DeleteChartVersionRequest struct {
	InstanceId   *string `json:"-"`
	ProjectName  *string `json:"-"`
	ChartName    *string `json:"-"`
	ChartVersion *string `json:"-"`
}
