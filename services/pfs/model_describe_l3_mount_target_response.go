package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeL3MountTargetResponse struct {
	bce.BaseResponse
	RequestId     *string `json:"requestId,omitempty"`
	Domain        *string `json:"domain,omitempty"`
	MountTargetId *string `json:"mountTargetId,omitempty"`
	Ovip          *string `json:"ovip,omitempty"`
	VpcId         *string `json:"vpcId,omitempty"`
	SubnetId      *string `json:"subnetId,omitempty"`
}
