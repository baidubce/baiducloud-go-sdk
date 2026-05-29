package cfs

type AccessRuleModel struct {
	AccessRuleId *int32  `json:"accessRuleId,omitempty"`
	Ip           *string `json:"ip,omitempty"`
	Mask         *int32  `json:"mask,omitempty"`
	Priority     *int32  `json:"priority,omitempty"`
	UserAccess   *string `json:"userAccess,omitempty"`
	WriteAccess  *string `json:"writeAccess,omitempty"`
}
