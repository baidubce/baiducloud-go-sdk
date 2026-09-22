package scs

type ManualBackupRequest struct {
	InstanceId *string `json:"-"`
	Comment    *string `json:"comment,omitempty"`
}
