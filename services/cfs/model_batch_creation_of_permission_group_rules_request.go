package cfs

type BatchCreationOfPermissionGroupRulesRequest struct {
	AgName      *string     `json:"ag_name,omitempty"`
	AccessRules []*RuleInfo `json:"access_rules,omitempty"`
}
