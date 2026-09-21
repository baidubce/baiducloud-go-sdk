package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListRedisBigKeyAnalysisTasksResponse struct {
	bce.BaseResponse
	DataCollectionTime *string             `json:"dataCollectionTime,omitempty"`
	ElementCountResult []*BigKeyResultInfo `json:"elementCountResult,omitempty"`
	MemoryResult       []*BigKeyResultInfo `json:"memoryResult,omitempty"`
}
