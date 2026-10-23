package pfs

type UpdateL3AccessRuleRequest struct {
	Action          *string `json:"-"`
	AccessGroupName *string `json:"accessGroupName,omitempty"`
	ZoneName        *string `json:"zoneName,omitempty"`
	AccessRuleId    *int32  `json:"accessRuleId,omitempty"`
	Ip              *string `json:"ip,omitempty"`
	Mask            *int32  `json:"mask,omitempty"`
	WriteAccess     *string `json:"writeAccess,omitempty"`
	UserAccess      *string `json:"userAccess,omitempty"`
	Priority        *int32  `json:"priority,omitempty"`
}
