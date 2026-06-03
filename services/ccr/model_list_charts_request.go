package ccr

type ListChartsRequest struct {
	InstanceId  *string `json:"-"`
	ProjectName *string `json:"-"`
	ChartName   *string `json:"-"`
	PageNo      *int32  `json:"-"`
	PageSize    *int32  `json:"-"`
}
