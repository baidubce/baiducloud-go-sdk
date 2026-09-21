package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetRedisBigKeyAnalysisResultResponse struct {
	bce.BaseResponse
	Tasks      []*BigKeyTask `json:"tasks,omitempty"`
	TotalCount *int32        `json:"totalCount,omitempty"`
}
