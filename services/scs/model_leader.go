package scs

type Leader struct {
	GroupName         *string  `json:"groupName,omitempty"`
	LeaderId          *string  `json:"leaderId,omitempty"`
	LeaderRegion      *string  `json:"leaderRegion,omitempty"`
	ClusterName       *string  `json:"clusterName,omitempty"`
	ClusterShowId     *string  `json:"clusterShowId,omitempty"`
	Region            *string  `json:"region,omitempty"`
	Status            *string  `json:"status,omitempty"`
	TotalCapacityInGB *float32 `json:"totalCapacityInGB,omitempty"`
	UsedCapacityInGB  *int32   `json:"usedCapacityInGB,omitempty"`
	ShardNum          *int32   `json:"shardNum,omitempty"`
	Flavor            *int32   `json:"flavor,omitempty"`
	QpsWrite          *int64   `json:"qpsWrite,omitempty"`
	QpsRead           *int64   `json:"qpsRead,omitempty"`
	StaleReadable     *bool    `json:"staleReadable,omitempty"`
	ForbidWrite       *int32   `json:"forbidWrite,omitempty"`
	AvailabilityZone  *string  `json:"availabilityZone,omitempty"`
	ExpiredTime       *string  `json:"expiredTime,omitempty"`
}
