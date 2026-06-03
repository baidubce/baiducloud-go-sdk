package ccr

type GetTagsScanOverviewRequest struct {
	InstanceId     *string `json:"-"`
	ProjectName    *string `json:"-"`
	RepositoryName *string `json:"-"`
	TagName        *string `json:"-"`
	PageNo         *int32  `json:"-"`
	PageSize       *int32  `json:"-"`
}
