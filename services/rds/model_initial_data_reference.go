package rds

type InitialDataReference struct {
	InstanceId    *string `json:"instanceId,omitempty"`
	ReferenceType *string `json:"referenceType,omitempty"`
	Datetime      *string `json:"datetime,omitempty"`
	SnapshotId    *string `json:"snapshotId,omitempty"`
}
