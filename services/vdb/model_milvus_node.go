package vdb

type MilvusNode struct {
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	ComponentType    *string `json:"componentType,omitempty"`
	FixedIp          *string `json:"fixedIp,omitempty"`
	FloatingIp       *string `json:"floatingIp,omitempty"`
	NodeId           *string `json:"nodeId,omitempty"`
	Port             *int32  `json:"port,omitempty"`
	Status           *string `json:"status,omitempty"`
}
