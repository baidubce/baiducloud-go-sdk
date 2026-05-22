package vpc

type CreateARegularSecurityGroupV2Request struct {
	ClientToken *string                   `json:"-"`
	Name        *string                   `json:"name,omitempty"`
	VpcId       *string                   `json:"vpcId,omitempty"`
	Desc        *string                   `json:"desc,omitempty"`
	Rules       []*SecurityGroupRuleModel `json:"rules,omitempty"`
	Tags        []*TagModel               `json:"tags,omitempty"`
}
