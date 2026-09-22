package scs

type CheckSyncGroupResultItem struct {
	MemberId           *string `json:"memberId,omitempty"`
	NoData             *bool   `json:"noData,omitempty"`
	Version            *bool   `json:"version,omitempty"`
	EngineVersion      *bool   `json:"engineVersion,omitempty"`
	ClusterStatus      *bool   `json:"clusterStatus,omitempty"`
	ShardNum           *bool   `json:"shardNum,omitempty"`
	ReplicationNum     *bool   `json:"replicationNum,omitempty"`
	Flavor             *bool   `json:"flavor,omitempty"`
	NotJoined          *bool   `json:"notJoined,omitempty"`
	NoSecurityGroup    *bool   `json:"noSecurityGroup,omitempty"`
	IsHitX1            *bool   `json:"isHitX1,omitempty"`
	IsAppendOnlyOn     *bool   `json:"isAppendOnlyOn,omitempty"`
	SamePasswd         *bool   `json:"samePasswd,omitempty"`
	HasSameHashTagConf *bool   `json:"hasSameHashTagConf,omitempty"`
	HasSetPwd          *bool   `json:"hasSetPwd,omitempty"`
}
