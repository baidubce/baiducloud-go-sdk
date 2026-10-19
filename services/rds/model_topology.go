package rds

type Topology struct {
	Master               []*string               `json:"master,omitempty"`
	ReadReplica          []*string               `json:"readReplica,omitempty"`
	Rdsproxy             []*string               `json:"rdsproxy,omitempty"`
	ReadReplicaIdMapping []*ReadReplicaIdMapping `json:"readReplicaIdMapping,omitempty"`
	MasterIdMapping      *MasterIdMapping        `json:"masterIdMapping,omitempty"`
	RdsproxyIdMapping    *ReadReplicaIdMapping   `json:"rdsproxyIdMapping,omitempty"`
	RdsproxyIdMappings   []*ReadReplicaIdMapping `json:"rdsproxyIdMappings,omitempty"`
}
