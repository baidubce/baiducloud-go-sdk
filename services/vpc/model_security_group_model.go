package vpc

type SecurityGroupModel struct {
	Id          *string                   `json:"id,omitempty"`
	Name        *string                   `json:"name,omitempty"`
	Desc        *string                   `json:"desc,omitempty"`
	VpcId       *string                   `json:"vpcId,omitempty"`
	CreatedTime *string                   `json:"createdTime,omitempty"`
	SgVersion   *int64                    `json:"sgVersion,omitempty"`
	Rules       []*SecurityGroupRuleModel `json:"rules,omitempty"`
	Tags        []*TagModel               `json:"tags,omitempty"`
}
