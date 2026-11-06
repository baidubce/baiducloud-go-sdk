package rds

type SlowLogDownloadDetailsRequest struct {
	InstanceId             *string `json:"-"`
	LogId                  *string `json:"-"`
	DownloadValidTimeInSec *int32  `json:"-"`
}
