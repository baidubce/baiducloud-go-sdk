package cfs

type MountTargetModel struct {
	AccessGroupName *string `json:"accessGroupName,omitempty"`
	Domain          *string `json:"domain,omitempty"`
	MountId         *string `json:"mountId,omitempty"`
	Ovip            *string `json:"ovip,omitempty"`
	SubnetId        *string `json:"subnetId,omitempty"`
	VpcId           *string `json:"vpcId,omitempty"`
}
