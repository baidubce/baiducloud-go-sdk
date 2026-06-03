package ccr

type DeleteTagsRequest struct {
	InstanceId     *string   `json:"-"`
	ProjectName    *string   `json:"-"`
	RepositoryName *string   `json:"-"`
	Items          []*string `json:"items,omitempty"`
}
