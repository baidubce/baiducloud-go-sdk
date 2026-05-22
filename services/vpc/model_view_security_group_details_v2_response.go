package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ViewSecurityGroupDetailsV2Response struct {
	bce.BaseResponse
	Id              *string                   `json:"id,omitempty"`
	Name            *string                   `json:"name,omitempty"`
	VpcId           *string                   `json:"vpcId,omitempty"`
	Desc            *string                   `json:"desc,omitempty"`
	CreatedTime     *string                   `json:"createdTime,omitempty"`
	SgVersion       *int64                    `json:"sgVersion,omitempty"`
	BindInstanceNum *int32                    `json:"bindInstanceNum,omitempty"`
	Rules           []*SecurityGroupRuleModel `json:"rules,omitempty"`
	Tags            []*TagModel               `json:"tags,omitempty"`
}
