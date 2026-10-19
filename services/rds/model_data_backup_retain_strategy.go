package rds

type DataBackupRetainStrategy struct {
	StartSeconds *int32 `json:"startSeconds,omitempty"`
	EndSeconds   *int32 `json:"endSeconds,omitempty"`
	RetainCount  *int32 `json:"retainCount,omitempty"`
	Precision    *int32 `json:"precision,omitempty"`
}
