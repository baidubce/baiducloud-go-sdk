package ccr

type UpdateRepositoryRequest struct {
	InstanceId     *string `json:"-"`
	ProjectName    *string `json:"-"`
	RepositoryName *string `json:"-"`
	Description    *string `json:"description,omitempty"`
}
