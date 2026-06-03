package ccr

type ListProjectsRequest struct {
	InstanceId  *string `json:"-"`
	ProjectName *string `json:"-"`
	PageNo      *int32  `json:"-"`
	PageSize    *int32  `json:"-"`
}
