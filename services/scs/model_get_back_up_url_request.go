package scs

type GetBackUpUrlRequest struct {
	InstanceId *string `json:"-"`
	BackupId   *int32  `json:"-"`
}
