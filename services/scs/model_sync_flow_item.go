package scs

type SyncFlowItem struct {
	TargetBLBIp         *string `json:"targetBLBIp,omitempty"`
	TargetBLBPort       *string `json:"targetBLBPort,omitempty"`
	TargetClusterShowId *string `json:"targetClusterShowId,omitempty"`
}
