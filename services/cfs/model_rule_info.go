package cfs

type RuleInfo struct {
	Ip          *string `json:"ip,omitempty"`
	Mask        *int32  `json:"mask,omitempty"`
	WriteAccess *string `json:"write_access,omitempty"`
	UserAccess  *string `json:"user_access,omitempty"`
	Priority    *int32  `json:"priority,omitempty"`
}
