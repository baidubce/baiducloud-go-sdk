package pfs

type L3ClientModel struct {
	ClientIp        *string `json:"clientIp,omitempty"`
	ClientId        *int64  `json:"clientId,omitempty"`
	MountId         *string `json:"mountId,omitempty"`
	MountTargetName *string `json:"mountTargetName,omitempty"`
	VpcId           *string `json:"vpcId,omitempty"`
	ZoneName        *string `json:"zoneName,omitempty"`
	ServerIp        *string `json:"serverIp,omitempty"`
	Protocol        *string `json:"protocol,omitempty"`
}
