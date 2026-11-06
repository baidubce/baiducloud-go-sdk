package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type RetrieveTheBinlogListResponse struct {
	bce.BaseResponse
	Binlogs []*Binlogs `json:"binlogs,omitempty"`
}
