package rds

type ErrorLogDownloadDetailsRequest struct {
	InstanceId             *string `json:"-"`
	LogId                  *string `json:"-"`
	DownloadValidTimeInSec *int32  `json:"-"`
}
