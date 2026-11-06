package rds

type QueryBackupListRequest struct {
	InstanceId *string `json:"-"`
	Marker     *string `json:"-"`
	MaxKeys    *int32  `json:"-"`
}
