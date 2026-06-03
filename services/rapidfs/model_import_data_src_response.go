package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ImportDataSrcResponse struct {
	bce.BaseResponse
	DataSrcId *string `json:"dataSrcId,omitempty"`
}
