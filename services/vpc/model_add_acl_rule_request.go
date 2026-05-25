package vpc

type AddAclRuleRequest struct {
	ClientToken *string           `json:"-"`
	AclRules    []*AclRuleRequest `json:"aclRules,omitempty"`
}
