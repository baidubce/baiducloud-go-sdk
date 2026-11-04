package pfs

type AccessRuleModel struct {
	Ip          *string `json:"ip,omitempty"`
	Mask        *int32  `json:"mask,omitempty"`
	WriteAccess *string `json:"writeAccess,omitempty"`
	UserAccess  *string `json:"userAccess,omitempty"`
	Priority    *int32  `json:"priority,omitempty"`
}
