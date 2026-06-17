package bcm

type AlarmPolicy struct {
	Id              *string         `json:"id,omitempty"`
	Name            *string         `json:"name,omitempty"`
	Scope           *string         `json:"scope,omitempty"`
	ResourceType    *string         `json:"resourceType,omitempty"`
	SubResourceType *string         `json:"subResourceType,omitempty"`
	Target          *AlarmTarget    `json:"target,omitempty"`
	Rules           []*AlarmRule    `json:"rules,omitempty"`
	Actions         []*PolicyAction `json:"actions,omitempty"`
	CreatedTime     *string         `json:"createdTime,omitempty"`
	UpdatedTime     *string         `json:"updatedTime,omitempty"`
}
