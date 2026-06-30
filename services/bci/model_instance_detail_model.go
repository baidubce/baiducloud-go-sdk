package bci

type InstanceDetailModel struct {
	Volume         *Volume                 `json:"volume,omitempty"`
	Containers     []*ContainerDetailModel `json:"containers,omitempty"`
	InitContainers []*ContainerDetailModel `json:"initContainers,omitempty"`
	SecurityGroups []*SecurityGroupModel   `json:"securityGroups,omitempty"`
	Vpc            *VpcModel               `json:"vpc,omitempty"`
	Subnet         *SubnetModel            `json:"subnet,omitempty"`
}
