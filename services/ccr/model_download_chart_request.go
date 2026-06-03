package ccr

type DownloadChartRequest struct {
	InstanceId  *string `json:"-"`
	ProjectName *string `json:"-"`
	Filename    *string `json:"-"`
}
