package aihc

type Pod struct {
	PodIP        *string `json:"PodIP,omitempty"`
	NodeName     *string `json:"nodeName,omitempty"`
	CreationAt   *string `json:"creationAt,omitempty"`
	Uid          *string `json:"uid,omitempty"`
	Name         *string `json:"name,omitempty"`
	Status       *string `json:"status,omitempty"`
	ReplicaType  *string `json:"replicaType,omitempty"`
	RestartCount *int32  `json:"restartCount,omitempty"`
	Envs         []*Env  `json:"envs,omitempty"`
	FinishedAt   *string `json:"finishedAt,omitempty"`
	Reason       *string `json:"reason,omitempty"`
}
