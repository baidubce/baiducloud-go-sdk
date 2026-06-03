package ccr

type DeleteTagRequest struct {
	InstanceId     *string `json:"-"`
	ProjectName    *string `json:"-"`
	RepositoryName *string `json:"-"`
	TagName        *string `json:"-"`
}
