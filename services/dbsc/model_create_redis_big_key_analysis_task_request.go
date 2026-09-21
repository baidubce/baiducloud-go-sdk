package dbsc

type CreateRedisBigKeyAnalysisTaskRequest struct {
	AppId      *string `json:"appId,omitempty"`
	ClusterId  *string `json:"clusterId,omitempty"`
	BackupType *int32  `json:"backupType,omitempty"`
	BackupId   *string `json:"backupId,omitempty"`
}
