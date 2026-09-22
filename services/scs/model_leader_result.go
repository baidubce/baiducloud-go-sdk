package scs

type LeaderResult struct {
	Version             *bool `json:"version,omitempty"`
	ClusterStatus       *bool `json:"clusterStatus,omitempty"`
	ReplicationNum      *bool `json:"replicationNum,omitempty"`
	Flavor              *bool `json:"flavor,omitempty"`
	Joined              *bool `json:"joined,omitempty"`
	NoPasswd            *bool `json:"noPasswd,omitempty"`
	NoSecurityGroup     *bool `json:"noSecurityGroup,omitempty"`
	IsHitX1             *bool `json:"isHitX1,omitempty"`
	NoTde               *bool `json:"noTde,omitempty"`
	HasSameHashTagConf  *bool `json:"hasSameHashTagConf,omitempty"`
	HasSetPwd           *bool `json:"hasSetPwd,omitempty"`
	IsNotCrossAzNearest *bool `json:"isNotCrossAzNearest,omitempty"`
}
