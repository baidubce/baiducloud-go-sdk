package vpc

type DeleteAclRuleRequest struct {
	AclRuleId   *string `json:"-"`
	ClientToken *string `json:"-"`
}
