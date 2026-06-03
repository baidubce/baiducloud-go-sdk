package ccr

type Repository struct {
	TagCount              *int32  `json:"tagCount,omitempty"`
	CreationTime          *string `json:"creationTime,omitempty"`
	Description           *string `json:"description,omitempty"`
	RepositoryName        *string `json:"repositoryName,omitempty"`
	PullCount             *int32  `json:"pullCount,omitempty"`
	UpdateTime            *string `json:"updateTime,omitempty"`
	RepositoryPath        *string `json:"repositoryPath,omitempty"`
	PrivateRepositoryPath *string `json:"privateRepositoryPath,omitempty"`
	ProjectName           *string `json:"projectName,omitempty"`
}
