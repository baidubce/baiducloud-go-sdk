package aihc

type Model struct {
	Id              *string `json:"id,omitempty"`
	Name            *string `json:"name,omitempty"`
	InitSource      *string `json:"initSource,omitempty"`
	LatestVersion   *string `json:"latestVersion,omitempty"`
	LatestVersionId *string `json:"latestVersionId,omitempty"`
	ModelFormat     *string `json:"modelFormat,omitempty"`
	Description     *string `json:"description,omitempty"`
	UpdatedAt       *string `json:"updatedAt,omitempty"`
	CreatedAt       *string `json:"createdAt,omitempty"`
	Owner           *string `json:"owner,omitempty"`
	OwnerName       *string `json:"ownerName,omitempty"`
	VisibilityScope *string `json:"visibilityScope,omitempty"`
}
