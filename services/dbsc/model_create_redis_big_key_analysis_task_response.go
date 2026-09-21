package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateRedisBigKeyAnalysisTaskResponse struct {
	bce.BaseResponse
	Id *string `json:"id,omitempty"`
}
