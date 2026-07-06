package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BatchStopTrainingTasksV2Response struct {
	bce.BaseResponse
	RequestId *string          `json:"requestId,omitempty"`
	Result    *StopBatchResult `json:"result,omitempty"`
}
