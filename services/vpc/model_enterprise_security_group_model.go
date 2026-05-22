package vpc

type EnterpriseSecurityGroupModel struct {
	Id          *string                             `json:"id,omitempty"`
	Name        *string                             `json:"name,omitempty"`
	Desc        *string                             `json:"desc,omitempty"`
	CreatedTime *string                             `json:"createdTime,omitempty"`
	UpdatedTime *string                             `json:"updatedTime,omitempty"`
	Rules       []*EnterpriseSecurityGroupRuleModel `json:"rules,omitempty"`
	Tags        []*TagModel                         `json:"tags,omitempty"`
}
