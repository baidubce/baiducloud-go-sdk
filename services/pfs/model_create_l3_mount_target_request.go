package pfs

type CreateL3MountTargetRequest struct {
	Action          *string `json:"-"`
	InstanceId      *string `json:"instanceId,omitempty"`
	VpcId           *string `json:"vpcId,omitempty"`
	SubnetId        *string `json:"subnetId,omitempty"`
	Name            *string `json:"name,omitempty"`
	AccessGroupName *string `json:"accessGroupName,omitempty"`
	Protocol        *string `json:"protocol,omitempty"`
	IpAddress       *string `json:"ipAddress,omitempty"`
}
