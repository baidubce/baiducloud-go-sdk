package ccr

type GetRepositoryRequest struct {
	InstanceId     *string `json:"-"`
	ProjectName    *string `json:"-"`
	RepositoryName *string `json:"-"`
}
