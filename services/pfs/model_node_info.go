package pfs

type NodeInfo struct {
	NodeId         *string   `json:"nodeId,omitempty"`
	NodeName       *string   `json:"nodeName,omitempty"`
	InternalIp     *string   `json:"internalIp,omitempty"`
	ZoneName       *string   `json:"zoneName,omitempty"`
	VpcId          *string   `json:"vpcId,omitempty"`
	VpcName        *string   `json:"vpcName,omitempty"`
	InstanceIdList []*string `json:"instanceIdList,omitempty"`
	NodeStatus     *string   `json:"nodeStatus,omitempty"`
	NodeType       *string   `json:"nodeType,omitempty"`
	MountStatus    *string   `json:"mountStatus,omitempty"`
	Passwd         *string   `json:"passwd,omitempty"`
	MtName         *string   `json:"mtName,omitempty"`
	MtId           *string   `json:"mtId,omitempty"`
	MtPath         *string   `json:"mtPath,omitempty"`
}
