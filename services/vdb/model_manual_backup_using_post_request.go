package vdb

type ManualBackupUsingPOSTRequest struct {
	InstanceId *string `json:"-"`
	EngineType *string `json:"-"`
	Comment    *string `json:"comment,omitempty"`
}
