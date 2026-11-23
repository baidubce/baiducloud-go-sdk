package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMysqlTableColumnsResponse struct {
	bce.BaseResponse
	Columns []*TableColumnItem `json:"columns,omitempty"`
}
