package scs

type FollowerResult struct {
	FollowerId          *string `json:"followerId,omitempty"`
	NoData              *bool   `json:"noData,omitempty"`
	Version             *bool   `json:"version,omitempty"`
	ClusterStatus       *bool   `json:"clusterStatus,omitempty"`
	ShardNum            *bool   `json:"shardNum,omitempty"`
	ReplicationNum      *bool   `json:"replicationNum,omitempty"`
	Flavor              *bool   `json:"flavor,omitempty"`
	Joined              *bool   `json:"joined,omitempty"`
	NoPasswd            *bool   `json:"noPasswd,omitempty"`
	NoSecurityGroup     *bool   `json:"noSecurityGroup,omitempty"`
	IsHitX1             *bool   `json:"isHitX1,omitempty"`
	NoTde               *bool   `json:"noTde,omitempty"`
	SamePasswd          *bool   `json:"samePasswd,omitempty"`
	HasSameHashTagConf  *bool   `json:"hasSameHashTagConf,omitempty"`
	HasSetPwd           *bool   `json:"hasSetPwd,omitempty"`
	IsNotCrossAzNearest *bool   `json:"isNotCrossAzNearest,omitempty"`
}
