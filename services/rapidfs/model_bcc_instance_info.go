package rapidfs

type BccInstanceInfo struct {
	Ip             *string `json:"ip,omitempty"`
	BccId          *string `json:"bccId,omitempty"`
	BccName        *string `json:"bccName,omitempty"`
	VpcId          *string `json:"vpcId,omitempty"`
	Zone           *string `json:"zone,omitempty"`
	BccSpec        *string `json:"bccSpec,omitempty"`
	Status         *string `json:"status,omitempty"`
	BsmAgentStatus *string `json:"bsmAgentStatus,omitempty"`
}
