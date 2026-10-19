package rds

type CheckStatusResponse struct {
	ClusterInstanceStatus *string `json:"clusterInstanceStatus,omitempty"`
	ClusterTopology       *string `json:"clusterTopology,omitempty"`
	ClusterMySQLIsAlived  *string `json:"clusterMySQLIsAlived,omitempty"`
	ClusterDelay          *string `json:"clusterDelay,omitempty"`
	LeaderReadOnly        *string `json:"leaderReadOnly,omitempty"`
}
