package ccr

type ListRepositoriesRequest struct {
	InstanceId     *string `json:"-"`
	ProjectName    *string `json:"-"`
	RepositoryName *string `json:"-"`
	PageNo         *int32  `json:"-"`
	PageSize       *int32  `json:"-"`
}
