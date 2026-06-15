package bcm

type DescribeAlarmTemplatesRequest struct {
	Name            *string `json:"name,omitempty"`
	Scope           *string `json:"scope,omitempty"`
	ResourceType    *string `json:"resourceType,omitempty"`
	SubResourceType *string `json:"subResourceType,omitempty"`
	Order           *string `json:"order,omitempty"`
	OrderBy         *string `json:"orderBy,omitempty"`
	PageNo          *int32  `json:"pageNo,omitempty"`
	PageSize        *int32  `json:"pageSize,omitempty"`
}
