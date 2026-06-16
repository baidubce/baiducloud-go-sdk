package bcc

type SecurityGroupModel struct {
	Id          *string                   `json:"id,omitempty"`
	Name        *string                   `json:"name,omitempty"`
	VpcId       *string                   `json:"vpcId,omitempty"`
	Desc        *string                   `json:"desc,omitempty"`
	CreatedTime *string                   `json:"createdTime,omitempty"`
	UpdatedTime *string                   `json:"updatedTime,omitempty"`
	SgVersion   *int32                    `json:"sgVersion,omitempty"`
	Rules       []*SecurityGroupRuleModel `json:"rules,omitempty"`
	Tags        []*Tag                    `json:"tags,omitempty"`
}
