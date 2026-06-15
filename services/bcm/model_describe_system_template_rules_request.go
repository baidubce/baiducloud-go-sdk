package bcm

type DescribeSystemTemplateRulesRequest struct {
	Scope           *string `json:"scope,omitempty"`
	ResourceType    *string `json:"resourceType,omitempty"`
	SubResourceType *string `json:"subResourceType,omitempty"`
	Source          *string `json:"source,omitempty"`
}
