package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryAclResponse struct {
	bce.BaseResponse
	VpcId     *string     `json:"vpcId,omitempty"`
	VpcName   *string     `json:"vpcName,omitempty"`
	VpcCidr   *string     `json:"vpcCidr,omitempty"`
	AclEntrys []*AclEntry `json:"aclEntrys,omitempty"`
}
