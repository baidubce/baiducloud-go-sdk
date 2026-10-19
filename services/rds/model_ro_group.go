package rds

type RoGroup struct {
	RoGroupId       *string       `json:"roGroupId,omitempty"`
	RoGroupName     *string       `json:"roGroupName,omitempty"`
	RoGroupEndpoint *InstanceDict `json:"roGroupEndpoint,omitempty"`
	VpcId           *string       `json:"vpcId,omitempty"`
	SubnetId        *string       `json:"subnetId,omitempty"`
	EipStatus       *string       `json:"eipStatus,omitempty"`
	RoGroupAppList  []*AppList    `json:"roGroupAppList,omitempty"`
	SimpleVpcVo     *SimpleVpcVo  `json:"SimpleVpcVo,omitempty"`
	SubnetVo        *SubnetVo     `json:"subnetVo,omitempty"`
	VpcVo           *SimpleVpcVo  `json:"vpcVo,omitempty"`
}
