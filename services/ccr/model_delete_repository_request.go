package ccr

type DeleteRepositoryRequest struct {
	InstanceId     *string `json:"-"`
	ProjectName    *string `json:"-"`
	RepositoryName *string `json:"-"`
}
