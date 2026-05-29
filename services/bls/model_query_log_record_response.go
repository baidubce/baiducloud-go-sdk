package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryLogRecordResponse struct {
	bce.BaseResponse
	NextMarker      *string          `json:"nextMarker,omitempty"`
	ResultSet       *ResultSet       `json:"resultSet,omitempty"`
	DatasetScanInfo *DatasetScanInfo `json:"datasetScanInfo,omitempty"`
}
