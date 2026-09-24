package scs

type GetBackUpUrlRequest struct {
	InstanceId *string `json:"-"`
	BackupId   *string `json:"-"`
}
