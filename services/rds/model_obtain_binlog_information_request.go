package rds

type ObtainBinlogInformationRequest struct {
	InstanceId             *string `json:"-"`
	LogId                  *string `json:"-"`
	DownloadValidTimeInSec *string `json:"-"`
}
