package pfs

type CreateL3AccessRuleRequest struct {
	Action          *string                 `json:"-"`
	AccessGroupName *string                 `json:"accessGroupName,omitempty"`
	ZoneName        *string                 `json:"zoneName,omitempty"`
	AccessRules     []*AccessRuleCreateInfo `json:"accessRules,omitempty"`
}
