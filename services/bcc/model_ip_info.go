package bcc

type IpInfo struct {
	PrivateIp       *string `json:"privateIp,omitempty"`
	Eip             *string `json:"eip,omitempty"`
	Primary         *string `json:"primary,omitempty"`
	EipId           *string `json:"eipId,omitempty"`
	EipAllocationId *string `json:"eipAllocationId,omitempty"`
	EipSize         *string `json:"eipSize,omitempty"`
	EipStatus       *string `json:"eipStatus,omitempty"`
	EipGroupId      *string `json:"eipGroupId,omitempty"`
	EipType         *string `json:"eipType,omitempty"`
}
