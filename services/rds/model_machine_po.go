package rds

type MachinePO struct {
	Id                    *int32  `json:"id,omitempty"`
	UpdatedTime           *string `json:"updatedTime,omitempty"`
	InstanceId            *string `json:"instanceId,omitempty"`
	MachineInstanceId     *string `json:"machineInstanceId,omitempty"`
	MachineInstanceName   *string `json:"machineInstanceName,omitempty"`
	MachineInstanceStatus *string `json:"machineInstanceStatus,omitempty"`
	Available             *bool   `json:"available,omitempty"`
	InstanceRole          *string `json:"instanceRole,omitempty"`
	MachineType           *string `json:"machineType,omitempty"`
	Azone                 *string `json:"azone,omitempty"`
}
