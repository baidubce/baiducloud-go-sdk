package vdb

type DeleteRecordUsingDELETERequest struct {
	InstanceId *string `json:"-"`
	EngineType *string `json:"-"`
	BatchId    *string `json:"-"`
	BackupId   *string `json:"-"`
}
