package rds

type GetPgLogDetailsRequest struct {
	InstanceId             *string `json:"-"`
	PglogId                *string `json:"-"`
	DownloadValidTimeInSec *int32  `json:"-"`
}
