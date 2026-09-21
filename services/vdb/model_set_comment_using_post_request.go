package vdb

type SetCommentUsingPOSTRequest struct {
	InstanceId *string `json:"-"`
	EngineType *string `json:"-"`
	BackupId   *string `json:"backupId,omitempty"`
	BatchId    *string `json:"batchId,omitempty"`
	Comment    *string `json:"comment,omitempty"`
}
