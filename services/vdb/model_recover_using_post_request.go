package vdb

type RecoverUsingPOSTRequest struct {
	InstanceId *string `json:"-"`
	EngineType *string `json:"-"`
	BackupId   *string `json:"backupId,omitempty"`
	BatchId    *string `json:"batchId,omitempty"`
	Confirmed  *bool   `json:"confirmed,omitempty"`
}
