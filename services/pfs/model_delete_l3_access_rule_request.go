package pfs

type DeleteL3AccessRuleRequest struct {
	Action          *string `json:"-"`
	AccessGroupName *string `json:"accessGroupName,omitempty"`
	ZoneName        *string `json:"zoneName,omitempty"`
	AccessRuleId    *int32  `json:"accessRuleId,omitempty"`
}
