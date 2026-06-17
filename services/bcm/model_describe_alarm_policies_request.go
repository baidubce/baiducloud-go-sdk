package bcm

type DescribeAlarmPoliciesRequest struct {
	PolicyName      *string `json:"policyName,omitempty"`
	PolicyId        *string `json:"policyId,omitempty"`
	Scope           *string `json:"scope,omitempty"`
	ResourceType    *string `json:"resourceType,omitempty"`
	Recursive       *bool   `json:"recursive,omitempty"`
	SubResourceType *string `json:"subResourceType,omitempty"`
	NotifyEnabled   *bool   `json:"notifyEnabled,omitempty"`
	BcmType         *string `json:"type,omitempty"`
	Order           *string `json:"order,omitempty"`
	OrderBy         *string `json:"orderBy,omitempty"`
	PageNo          *int32  `json:"pageNo,omitempty"`
	PageSize        *int32  `json:"pageSize,omitempty"`
}
