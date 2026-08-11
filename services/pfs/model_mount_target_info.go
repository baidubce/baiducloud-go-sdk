package pfs

type MountTargetInfo struct {
	Domain        *string `json:"domain,omitempty"`
	MountTargetId *string `json:"mountTargetId,omitempty"`
	Ovip          *string `json:"ovip,omitempty"`
	VpcId         *string `json:"vpcId,omitempty"`
	SubnetId      *string `json:"subnetId,omitempty"`
}
