package ccr

type ListChartVersionsRequest struct {
	InstanceId   *string `json:"-"`
	ProjectName  *string `json:"-"`
	ChartName    *string `json:"-"`
	ChartVersion *string `json:"-"`
	PageNo       *int32  `json:"-"`
	PageSize     *int32  `json:"-"`
}
