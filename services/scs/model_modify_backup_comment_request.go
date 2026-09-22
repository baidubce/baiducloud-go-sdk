package scs

type ModifyBackupCommentRequest struct {
	InstanceId *string `json:"-"`
	BatchId    *string `json:"-"`
	Comment    *string `json:"comment,omitempty"`
}
