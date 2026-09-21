package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListRedisSlowLogsResponse struct {
	bce.BaseResponse
	AppId                    *string           `json:"appId,omitempty"`
	NodeId                   *string           `json:"nodeId,omitempty"`
	RequestId                *string           `json:"requestId,omitempty"`
	Start                    *string           `json:"start,omitempty"`
	End                      *string           `json:"end,omitempty"`
	TotalRecordCount         *int32            `json:"totalRecordCount,omitempty"`
	MaxRecordCountPerPage    *int32            `json:"maxRecordCountPerPage,omitempty"`
	PageNumber               *int32            `json:"pageNumber,omitempty"`
	RecordCountInCurrentPage *int32            `json:"recordCountInCurrentPage,omitempty"`
	Records                  []*SCSSlowLogInfo `json:"records,omitempty"`
}
