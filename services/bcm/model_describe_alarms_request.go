package bcm

type DescribeAlarmsRequest struct {
	StartTime    *string `json:"startTime,omitempty"`
	EndTime      *string `json:"endTime,omitempty"`
	PolicyName   *string `json:"policyName,omitempty"`
	Scope        *string `json:"scope,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	State        *string `json:"state,omitempty"`
	BcmType      *string `json:"type,omitempty"`
	Order        *string `json:"order,omitempty"`
	OrderBy      *string `json:"orderBy,omitempty"`
	PageNo       *int32  `json:"pageNo,omitempty"`
	PageSize     *int32  `json:"pageSize,omitempty"`
}
