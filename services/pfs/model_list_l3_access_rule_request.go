package pfs

type ListL3AccessRuleRequest struct {
	Action          *string `json:"-"`
	AccessGroupName *string `json:"accessGroupName,omitempty"`
	ZoneName        *string `json:"zoneName,omitempty"`
	PageNo          *int32  `json:"pageNo,omitempty"`
	PageSize        *int32  `json:"pageSize,omitempty"`
	OrderBy         *string `json:"orderBy,omitempty"`
	Order           *string `json:"order,omitempty"`
	IpFilter        *string `json:"ipFilter,omitempty"`
}
