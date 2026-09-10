package cce

type MachineSpecStatus struct {
	MachineSpec *string `json:"machineSpec,omitempty"`
	Status      *string `json:"status,omitempty"`
}
