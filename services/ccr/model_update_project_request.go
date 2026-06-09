package ccr

type UpdateProjectRequest struct {
	InstanceId  *string `json:"-"`
	ProjectName *string `json:"-"`
	AutoScan    *string `json:"autoScan,omitempty"`
	Public      *string `json:"public,omitempty"`
}
