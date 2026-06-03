package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeAllocatableDataSrcQuotaResponse struct {
	bce.BaseResponse
	DataSrcQuotaInfo *DataSrcQuotaInfo `json:"dataSrcQuotaInfo,omitempty"`
}
