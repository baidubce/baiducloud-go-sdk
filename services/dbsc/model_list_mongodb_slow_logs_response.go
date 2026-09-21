package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListMongodbSlowLogsResponse struct {
	bce.BaseResponse
	Logs       []*MongoDBSlowLogDetail `json:"logs,omitempty"`
	TotalCount *int32                  `json:"totalCount,omitempty"`
}
