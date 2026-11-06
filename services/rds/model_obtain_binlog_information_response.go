package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ObtainBinlogInformationResponse struct {
	bce.BaseResponse
	Binlog *Binlog `json:"binlog,omitempty"`
}
