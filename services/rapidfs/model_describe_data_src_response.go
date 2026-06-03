package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeDataSrcResponse struct {
	bce.BaseResponse
	DataSrcInfo *DataSrcInfo `json:"dataSrcInfo,omitempty"`
}
