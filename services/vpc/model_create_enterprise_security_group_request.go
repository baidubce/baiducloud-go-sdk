package vpc

type CreateEnterpriseSecurityGroupRequest struct {
	ClientToken *string                             `json:"-"`
	Name        *string                             `json:"name,omitempty"`
	Desc        *string                             `json:"desc,omitempty"`
	Rules       []*EnterpriseSecurityGroupRuleModel `json:"rules,omitempty"`
	Tags        []*TagModel                         `json:"tags,omitempty"`
}
