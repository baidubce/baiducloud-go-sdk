package ccr

type CreateProjectRequest struct {
	InstanceId  *string `json:"-"`
	ProjectName *string `json:"projectName,omitempty"`
	Public      *string `json:"public,omitempty"`
}
