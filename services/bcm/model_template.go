package bcm

type Template struct {
	Scope           *string      `json:"scope,omitempty"`
	ResourceType    *string      `json:"resourceType,omitempty"`
	SubResourceType *string      `json:"subResourceType,omitempty"`
	Name            *string      `json:"name,omitempty"`
	Comment         *string      `json:"comment,omitempty"`
	Rules           []*AlarmRule `json:"rules,omitempty"`
}
