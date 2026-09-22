package scs

type CheckList struct {
	ClusterInstanceStatus *string `json:"clusterInstanceStatus,omitempty"`
	ClusterTopology       *string `json:"clusterTopology,omitempty"`
	ClusterRedisIsAlived  *string `json:"clusterRedisIsAlived,omitempty"`
	ClusterDelay          *string `json:"clusterDelay,omitempty"`
	LeaderReadOnly        *string `json:"leaderReadOnly,omitempty"`
}
