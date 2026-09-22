package scs

type SetClusterAsSlaveRequest struct {
	InstanceId   *string `json:"-"`
	MasterDomain *string `json:"masterDomain,omitempty"`
	MasterPort   *int32  `json:"masterPort,omitempty"`
}
