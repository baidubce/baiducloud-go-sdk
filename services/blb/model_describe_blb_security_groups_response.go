package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeBlbSecurityGroupsResponse struct {
	bce.BaseResponse
	BlbSecurityGroups []*BlbSecurityGroupModel `json:"blbSecurityGroups,omitempty"`
}
