package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeBlbEnterpriseSecurityGroupsResponse struct {
	bce.BaseResponse
	EnterpriseSecurityGroups []*BlbEnterpriseSecurityGroupModel `json:"enterpriseSecurityGroups,omitempty"`
}
