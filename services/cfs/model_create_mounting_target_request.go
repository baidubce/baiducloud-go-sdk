package cfs

type CreateMountingTargetRequest struct {
	FsId            *string `json:"-"`
	SubnetId        *string `json:"subnetId,omitempty"`
	VpcId           *string `json:"vpcId,omitempty"`
	AccessGroupName *string `json:"accessGroupName,omitempty"`
	Address         *string `json:"address,omitempty"`
}
